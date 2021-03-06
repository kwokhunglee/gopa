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

package filter

import (
	log "github.com/cihub/seelog"
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
)



type LeveldbFilter struct{
	persistFileName string
	filter *leveldb.DB
	l sync.RWMutex
}


func (filter *LeveldbFilter) Open(fileName string) error{
	filter.l.Lock()
	defer filter.l.Unlock()

	filter.persistFileName=fileName
	db, err := leveldb.OpenFile(fileName, nil)
	filter.filter = db

	if err != nil {
		log.Error("leveldb:", fileName,", ", err)
		return err
	}

	log.Debug("leveldb successfully reloaded:", fileName)
	return nil
}

func (this *LeveldbFilter) Close() error{
	this.l.Lock()
	defer this.l.Unlock()
	log.Debug("start persist leveldb, file:",this.persistFileName)

	err := this.filter.Close()
	if err != nil {
		log.Error("leveldb:", this.persistFileName, err)
	}
	return err

}

func (filter *LeveldbFilter) Exists(key []byte) bool{
	filter.l.RLock()
	defer filter.l.RUnlock()
	value,_  :=  filter.filter.Get(key,nil)
	if value != nil {
		return true
	}
	return false
}

func (filter *LeveldbFilter) Add(key []byte) error{
	filter.l.Lock()
	defer filter.l.Unlock()
	filter.filter.Put(key,[]byte(""),nil)
	return nil
}
