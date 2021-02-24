package model

import (
	"database/sql"
	"encoding/json"
	"eventAnalysis/pkg/db"
	"k8s.io/klog"
	"log"
)

type PullImageInfo struct {
	EventName     string `json:"event_name"`
	ImageName     string `json:"image_name"`
	PullimageTime int32  `json:"pullimage_time"`
}

func (c *PullImageInfo) ToJsonString() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type PullImageInfos []PullImageInfo

func (c *PullImageInfos) ToJsonString() string {
	b, _ := json.Marshal(c)
	return string(b)
}
func UnmarshalPullImageInfo(jsonbody string) PullImageInfos {
	var obj PullImageInfos
	klog.Info(jsonbody)
	if jsonbody != "" {
		v := []byte(jsonbody)
		err := json.Unmarshal(v, &obj)
		if err != nil {
			klog.Info(err)
		}
	} else {
		klog.Error("jsonbody is null")
	}
	return obj
}

//展示数据
func Show() *PullImageInfos {
	db.InitDB()
	var pullImageInfo PullImageInfo
	var pullImageInfos PullImageInfos
	rows, _ := db.DB.Query("SELECT * FROM event_info")
	for rows.Next() {
		err := rows.Scan(&pullImageInfo.EventName, &pullImageInfo.ImageName, &pullImageInfo.PullimageTime)
		switch {
		case err == sql.ErrNoRows:
			log.Printf("No user with that ID.")
		case err != nil:
			log.Fatal(err)
		}
		pullImageInfos = append(pullImageInfos, pullImageInfo)
	}
	return &pullImageInfos
}

//写数据
func InsertEvent(pullImageInfo *PullImageInfo) {
	if pullImageInfo != nil {
		klog.Info(pullImageInfo.ToJsonString())
		insertsql := "INSERT INTO event_info(event_name,image_name,pullimage_time) VALUES (?,?,?)"
		_, err := db.DB.Exec(insertsql, pullImageInfo.EventName, pullImageInfo.ImageName, pullImageInfo.PullimageTime)
		if err != nil {
			if err != nil {
				klog.Error("insert failed, err:%v\n", err)
				return
			}
		}
	} else {
		panic("post data is nil")
	}
}
