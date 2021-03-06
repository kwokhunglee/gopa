/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package url_checker

import (
	log "github.com/cihub/seelog"
	. "github.com/medcl/gopa/core/env"
	. "github.com/medcl/gopa/core/filter"
	"github.com/medcl/gopa/core/stats"
	"path"
	"time"
)

var quitChannel chan bool
var started = false
var filter = LeveldbFilter{}
var filterFileName = "filters/url_fetched"

func Start(env *Env) {
	if started {
		log.Error("url checker is already started, please stop it first.")
		return
	}
	quitChannel = make(chan bool)

	filter.Open(path.Join(env.RuntimeConfig.PathConfig.Data, filterFileName))

	go runCheckerGo(env, &quitChannel)
	started = true
}

func runCheckerGo(env *Env, quitC *chan bool) {

	go func() {
		for {
			startTime := time.Now()
			if !started {
				return
			}
			log.Trace("waiting url to check")

			url, err := env.Channels.PopUrlToCheck()

			stats.Increment("checker.url", "walk")
			if err != nil {
				continue
			}
			log.Trace("cheking url:", string(url.Url))

			//TODO 统一 url 格式 , url 目前可能是相对路径
			//checking
			if filter.Exists([]byte(url.Url)) {
				log.Debug("url already pushed to fetch queue, ignore :", string(url.Url))
				continue
			}

			//add to filter
			filter.Add([]byte(url.Url))

			//send to disk queue
			env.Channels.PushUrlToFetch(url)

			stats.Increment("checker.url", "get_valid_seed")

			log.Debugf("send url: %s ,depth: %d to  fetch queue", string(url.Url), url.Depth)
			elapsedTime := time.Now().Sub(startTime)
			stats.Timing("checker.url", "time", elapsedTime.Nanoseconds())
		}
	}()

	log.Trace("url checker success started")

	<-*quitC

	log.Info("url checker success stoped")
}

func Stop() error {
	if started {
		log.Debug("start shutting down url checker")

		filter.Close()

		quitChannel <- true

		started = false
	} else {
		log.Error("url checker is not started")
	}

	return nil
}
