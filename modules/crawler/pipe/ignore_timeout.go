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

package pipe

import (
	. "github.com/medcl/gopa/core/pipeline"
	log "github.com/cihub/seelog"
	"github.com/medcl/gopa/core/stats"
)

func (this IgnoreTimeoutJoint) Name() string {
	return "ignore_timeout"
}

type IgnoreTimeoutJoint struct {
	IgnoreTimeoutAfterCount int64
}

func (this IgnoreTimeoutJoint) Process(context *Context) (*Context, error) {


	host:=context.MustGetString(CONTEXT_HOST)
	timeoutCount:=stats.Stat(host,stats.STATS_FETCH_TIMEOUT_COUNT)
	if(timeoutCount>this.IgnoreTimeoutAfterCount){
		context.Break()
		log.Warnf("hit timeout host, %s , ignore after,%d ",host,timeoutCount)
	}
	return context, nil
}
