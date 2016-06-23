package main

import (
	"encoding/json"
	"fmt"
	"github.com/llitfkitfk/cirkol/pkg/result"
	"github.com/parnurzeal/gorequest"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"sync"
	"time"
	"reflect"
	"github.com/llitfkitfk/cirkol/pkg/models"
)

var tokenCh chan string
var popchan chan string

type Book struct {
	Title       string `json:"title"`
	CoTitle     string `json:"co_title"`
	Cover       string `json:"cover"`
	QRCode      string `json:"qr_code"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Year        string `json:"year"`
	Pages       string `json:"pages"`
	Language    string `json:"language"`
	FileSize    string `json:"file_size"`
	FileFormat  string `json:"file_format"`
	Download    string `json:"download"`
	Link        string `json:"link"`
}

func main() {
	//StartTokenGen()
	//URLTest()
	//RegexpDemo()
	//GoRoutine()
	//WaitGroupDemo()
	//RequestDemo()
	//MongoQuery()

	//TimeParse()
	//Decode()
	//jsonParse()
	var rawData models.FBRawProfile
	rawData.Name = "fb"

	reflectDemo(rawData)

}
func reflectDemo(v interface{}) {
	log.Println(reflect.TypeOf(v).Name())



}
func jsonParse() {

	src := `{"maxPage": 880, "page": 1, "url": "\/page\/json?containerid=1005051266321801_-_WEIBO_SECOND_PROFILE_WEIBO", "previous_cursor": "", "next_cursor": "", "loadMore": true, "card_group": [{
        "card_type": 9,
        "mblog": {
            "created_at": "\u4eca\u5929 12:26",
            "id": 3988419455558502,
            "mid": "3988419455558502",
            "idstr": "3988419455558502",
            "text": "\u6211\u68a6\u60f3\u6709\u4e00\u5929\uff0c\u5b57\u5178\u91cc\u518d\u6ca1\u6709\u201c\u96be\u6c11\u201d\u8fd9\u4e2a\u8bcd\uff0c\u6211\u4eec\u7684\u4e16\u754c\u518d\u6ca1\u6709\u6218\u706b\u548c\u9965\u997f\uff0c\u4f60\u4eec\u7684\u773c\u775b\u53ea\u770b\u5f97\u5230\u5feb\u4e50\uff0c\u518d\u6ca1\u6709\u6050\u60e7\u548c\u60b2\u4f24\u3002<i class=\"face face_3 icon_12\">[\u9c9c\u82b1]<\/i><a class='k' href='\/k\/ \u4e16\u754c\u96be\u6c11\u65e5?from=feed'># \u4e16\u754c\u96be\u6c11\u65e5#<\/a> <a href='\/n\/\u8054\u5408\u56fd\u96be\u6c11\u7f72'>@\u8054\u5408\u56fd\u96be\u6c11\u7f72<\/a>",
            "textLength": 145,
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": ["4b7a8989jw1f51jvxo4tjj20m80et3z8", "4b7a8989jw1f51jvxblnxj20zk0np40i", "4b7a8989jw1f51jvxpubcj20no0zkabz", "4b7a8989jw1f51jvxyf5yj20zk0nomzh", "4b7a8989jw1f51jvy9qdrj20zk0nqq5d", "4b7a8989jw1f51jvydyj6j20zk0nqtdc", "4b7a8989jw1f51jwj0kf3j215o0rs10y", "4b7a8989jw1f51jvz3c5qj21jk1111kx", "4b7a8989jw1f51jvzkpj9j20zk0npjzk"],
            "thumbnail_pic": "http:\/\/ww4.sinaimg.cn\/thumbnail\/4b7a8989jw1f51jvxo4tjj20m80et3z8.jpg",
            "bmiddle_pic": "http:\/\/ww4.sinaimg.cn\/bmiddle\/4b7a8989jw1f51jvxo4tjj20m80et3z8.jpg",
            "original_pic": "http:\/\/ww4.sinaimg.cn\/large\/4b7a8989jw1f51jvxo4tjj20m80et3z8.jpg",
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "picStatus": "0:1,1:1,2:1,3:1,4:1,5:1,6:1,7:1,8:1",
            "reposts_count": 907,
            "comments_count": 902,
            "attitudes_count": 17846,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 4294967300,
            "page_type": 32,
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "userType": 0,
            "extend_info": {
                "weibo_camera": {
                    "c": ["27235472_27235473_25857858_27235475_27366580_27366581_27235478_27235479_27235480"]
                }
            },
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "topic_struct": [{
                "topic_title": " \u4e16\u754c\u96be\u6c11\u65e5",
                "topic_url": "sinaweibo:\/\/pageinfo?containerid=100808bc5be94794751d15a8c3d671d5536357&pageid=100808bc5be94794751d15a8c3d671d5536357&extparam=%E4%!B(MISSING)8%E7%8C%!E(MISSING)9%!A(MISSING)%!B(MISSING)E%!E(MISSING)6%!B(MISSING)0%E6%A5"
            }],
            "created_timestamp": 1466396761,
            "bid": "DB6c1nk1g",
            "pics": [{
                "pid": "4b7a8989jw1f51jvxo4tjj20m80et3z8",
                "url": "http:\/\/ww4.sinaimg.cn\/thumb180\/4b7a8989jw1f51jvxo4tjj20m80et3z8.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 36957
                }
            }, {
                "pid": "4b7a8989jw1f51jvxblnxj20zk0np40i",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f51jvxblnxj20zk0np40i.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 84061
                }
            }, {
                "pid": "4b7a8989jw1f51jvxpubcj20no0zkabz",
                "url": "http:\/\/ww4.sinaimg.cn\/thumb180\/4b7a8989jw1f51jvxpubcj20no0zkabz.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 81175
                }
            }, {
                "pid": "4b7a8989jw1f51jvxyf5yj20zk0nomzh",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f51jvxyf5yj20zk0nomzh.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 95883
                }
            }, {
                "pid": "4b7a8989jw1f51jvy9qdrj20zk0nqq5d",
                "url": "http:\/\/ww3.sinaimg.cn\/thumb180\/4b7a8989jw1f51jvy9qdrj20zk0nqq5d.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 100072
                }
            }, {
                "pid": "4b7a8989jw1f51jvydyj6j20zk0nqtdc",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f51jvydyj6j20zk0nqtdc.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 181061
                }
            }, {
                "pid": "4b7a8989jw1f51jwj0kf3j215o0rs10y",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f51jwj0kf3j215o0rs10y.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 313344
                }
            }, {
                "pid": "4b7a8989jw1f51jvz3c5qj21jk1111kx",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f51jvz3c5qj21jk1111kx.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1048576
                }
            }, {
                "pid": "4b7a8989jw1f51jvzkpj9j20zk0npjzk",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f51jvzkpj9j20zk0npjzk.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 311854
                }
            }],
            "like_count": 17846,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-19 19:32",
            "id": 3988164370324127,
            "mid": "3988164370324127",
            "idstr": "3988164370324127",
            "text": "\u4f60\u6293\u7684\u8fd9\u53ea\u8001\u9f20\u4e2d\u770b\u4e0d\u4e2d\u5403\uff0c\u7c91\u7c91\u5f53\u7136\u4e0d\u5f00\u5fc3\u2026<i class=\"face face_1 icon_20\">[doge]<\/i>[\u8001\u9f20]",
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": [],
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "retweeted_status": {
                "created_at": "06-19 11:42",
                "id": 3988046023951109,
                "mid": "3988046023951109",
                "idstr": "3988046023951109",
                "text": "\u4eca\u5929\u662f<a class='k' href='\/k\/\u7236\u4eb2\u8282?from=feed'>#\u7236\u4eb2\u8282#<\/a> \uff0c\u7279\u610f\u6293\u4e86\u53ea\u5927\u8001\u9f20\u5b5d\u656c\u7238\u7238\uff0c\u4e0d\u8fc7\u4ed6\u6536\u5230\u4e86\u597d\u50cf\u5e76\u4e0d\u592a\u5f00\u5fc3\uff0c\u70b9\u89e3\uff1f\u6025\uff0c\u5728\u7ebf\u7b49[\u50bb\u773c]<a href='\/n\/\u59da\u6668'>@\u59da\u6668<\/a>",
                "textLength": 98,
                "source_allowclick": 0,
                "source_type": 1,
                "source": "iPhone 6s",
                "favorited": false,
                "pic_ids": ["ac8858d9jw1f50dvbzcujj20ou0x448y"],
                "thumbnail_pic": "http:\/\/ww3.sinaimg.cn\/thumbnail\/ac8858d9jw1f50dvbzcujj20ou0x448y.jpg",
                "bmiddle_pic": "http:\/\/ww3.sinaimg.cn\/bmiddle\/ac8858d9jw1f50dvbzcujj20ou0x448y.jpg",
                "original_pic": "http:\/\/ww3.sinaimg.cn\/large\/ac8858d9jw1f50dvbzcujj20ou0x448y.jpg",
                "user": {
                    "id": 2894616793,
                    "screen_name": "\u66f9\u516b\u987f",
                    "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.0.512.512.180\/ac8858d9jw8etia9r1pcnj20e80e8aar.jpg",
                    "profile_url": "\/u\/2894616793",
                    "statuses_count": 206,
                    "verified": false,
                    "verified_reason": "",
                    "description": "\u59da\u6668\u5de5\u4f5c\u5ba4\u603b\u88c1  \u6f14\u5458\u59da\u6668\u7684\u55b5\u4e3b\u5b50",
                    "remark": "",
                    "verified_type": -1,
                    "gender": "m",
                    "mbtype": 0,
                    "h5icon": {
                        "main": 0,
                        "other": []
                    },
                    "ismember": 0,
                    "valid": null,
                    "fansNum": 12363,
                    "follow_me": false,
                    "following": false
                },
                "picStatus": "0:1",
                "reposts_count": 66,
                "comments_count": 53,
                "attitudes_count": 328,
                "isLongText": false,
                "mlevel": 0,
                "visible": {
                    "type": 0,
                    "list_id": 0
                },
                "biz_feature": 4294967300,
                "page_type": 32,
                "hot_weibo_tags": [],
                "text_tag_tips": [],
                "userType": 0,
                "positive_recom_flag": 0,
                "page_info": {
                    "page_id": "1008080ee883fb848389af3c49b2e5e8cc9042",
                    "type": 0,
                    "page_title": "#\u7236\u4eb2\u8282#",
                    "page_url": "http:\/\/m.weibo.cn\/p\/index?containerid=1008080ee883fb848389af3c49b2e5e8cc9042",
                    "page_pic": "http:\/\/ww3.sinaimg.cn\/thumb300\/719d0099jw1etcpnwcsp5j2050050t8s.jpg",
                    "page_desc": "\u7236\u4eb2\u8282\u6765\u4e34\uff0c\u5fae\u535a\u8282\u65e5\u795d\u6240\u6709\u7238\u7238\uff0c\u7236\u4eb2\u8282\u5feb\u4e50\uff01",
                    "object_type": "topic",
                    "tips": "10016\u4eba\u5173\u6ce8",
                    "object_id": "1022:1008080ee883fb848389af3c49b2e5e8cc9042",
                    "actionlog": {
                        "oid": "1008080ee883fb848389af3c49b2e5e8cc9042",
                        "act_code": 300,
                        "act_type": 1,
                        "ext": "mid:3988164370324127|rid:"
                    },
                    "content1": "",
                    "content2": ""
                },
                "created_timestamp": 1466307728,
                "bid": "DAWtIgzRz",
                "pics": [{
                    "pid": "ac8858d9jw1f50dvbzcujj20ou0x448y",
                    "url": "http:\/\/ww3.sinaimg.cn\/wap180\/ac8858d9jw1f50dvbzcujj20ou0x448y.jpg",
                    "size": "wap180",
                    "geo": {
                        "width": 135,
                        "height": 180,
                        "croped": false,
                        "byte": 395357
                    }
                }],
                "like_count": 328,
                "attitudes_status": 0
            },
            "reposts_count": 60,
            "comments_count": 570,
            "attitudes_count": 17851,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 0,
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "userType": 0,
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "topic_struct": [{
                "topic_title": "\u7236\u4eb2\u8282",
                "topic_url": "sinaweibo:\/\/pageinfo?containerid=1008080ee883fb848389af3c49b2e5e8cc9042&pageid=1008080ee883fb848389af3c49b2e5e8cc9042&extparam=%!E(MISSING)7%B6%!E(MISSING)4%!B(MISSING)A%!B(MISSING)2%!E(MISSING)8%!A(MISSING)%!"
            }],
            "created_timestamp": 1466335945,
            "bid": "DAZyB1mjR",
            "like_count": 17851,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-19 12:53",
            "id": 3988063988295621,
            "mid": "3988063988295621",
            "idstr": "3988063988295621",
            "text": "\u4e24\u4f4d\u597d\u7238\u7238\uff0c\u628a\u5a03\u62c9\u626f\u5927\u662f\u4ef6\u591a\u4e48\u4e0d\u5bb9\u6613\u7684\u4e8b\u60c5\u554a\uff01\u8001\u59da\u66f9\u8001\uff0c\u4f60\u4eec\u8f9b\u82e6\uff0c\u7236\u4eb2\u8282\u5feb\u4e50\uff01[\u5f20\u5634]<a class='k' href='\/k\/\u8c22\u8c22\u4f60\u7684\u7231?from=feed'>#\u8c22\u8c22\u4f60\u7684\u7231#<\/a>",
            "textLength": 94,
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": ["4b7a8989jw1f50e1n5vtyj20ku0kuwuo", "4b7a8989jw1f50dy05gxrj21r91e9kjq", "4b7a8989jw1f50ea8r4brj21w01w0e84", "4b7a8989jw1f50fvq2tcoj21a31dwnpf"],
            "thumbnail_pic": "http:\/\/ww1.sinaimg.cn\/thumbnail\/4b7a8989jw1f50e1n5vtyj20ku0kuwuo.jpg",
            "bmiddle_pic": "http:\/\/ww1.sinaimg.cn\/bmiddle\/4b7a8989jw1f50e1n5vtyj20ku0kuwuo.jpg",
            "original_pic": "http:\/\/ww1.sinaimg.cn\/large\/4b7a8989jw1f50e1n5vtyj20ku0kuwuo.jpg",
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "stickerID": "0:10600",
            "picStatus": "0:1,1:1,2:1,3:1",
            "reposts_count": 210,
            "comments_count": 1040,
            "attitudes_count": 42430,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 4294967300,
            "page_type": 32,
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "userType": 0,
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "topic_struct": [{
                "topic_title": "\u8c22\u8c22\u4f60\u7684\u7231",
                "topic_url": "sinaweibo:\/\/pageinfo?containerid=100808fa05f3683d2590bb18c6cd105d9b1493&pageid=100808fa05f3683d2590bb18c6cd105d9b1493&extparam=%!E(MISSING)8%!B(MISSING)0%!A(MISSING)2%!E(MISSING)8%!B(MISSING)0%!A(MISSING)2%!E(MISSING)4%!B(MISSING)D%!A(MISSING)0%!E(MISSING)7%!A(MISSING)%E7%B1"
            }],
            "page_info": {
                "page_id": "100808fa05f3683d2590bb18c6cd105d9b1493",
                "type": 0,
                "page_title": "#\u8c22\u8c22\u4f60\u7684\u7231#",
                "page_url": "http:\/\/m.weibo.cn\/p\/index?containerid=100808fa05f3683d2590bb18c6cd105d9b1493",
                "page_pic": "http:\/\/ww1.sinaimg.cn\/thumb150\/71fb780djw1eaj6auhkxyj20500500sx.jpg",
                "page_desc": "\u6211\u8981\u8c22\u8c22\u4f60\u7684\u7231\uff0c\u662f\u4f60\u7684\u7231\u8ba9\u6211\u5728\u8fd9\u4e2a\u51ac\u5929\u4e0d\u5b64\u5355\u3001\u5728\u594b\u6597\u7684\u8def\u4e0a\u4e0d\u6015\u8f9b\u82e6\u3001\u5728\u6210\u957f\u7684\u9053\u8def\u4e0a\u6709\u6240\u4f9d\u8d56\u3002\u6211\u8981\u8c22\u8c22\u4f60\u5e26\u6211\u770b\u8fc7\u7684\u6240\u6709\u98ce\u666f\uff0c\u8c22\u8c22\u4f60\u548c\u6211\u7684\u65c5\u9014\u6545\u4e8b\u3002",
                "object_type": "topic",
                "tips": "214\u4eba\u5173\u6ce8",
                "object_id": "1022:100808fa05f3683d2590bb18c6cd105d9b1493",
                "actionlog": {
                    "oid": "100808fa05f3683d2590bb18c6cd105d9b1493",
                    "act_code": 300,
                    "act_type": 1,
                    "ext": "mid:3988063988295621|rid:"
                },
                "content1": "",
                "content2": ""
            },
            "created_timestamp": 1466312012,
            "bid": "DAWWGyO4l",
            "pics": [{
                "pid": "4b7a8989jw1f50e1n5vtyj20ku0kuwuo",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f50e1n5vtyj20ku0kuwuo.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 607138
                }
            }, {
                "pid": "4b7a8989jw1f50dy05gxrj21r91e9kjq",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f50dy05gxrj21r91e9kjq.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 6862848
                }
            }, {
                "pid": "4b7a8989jw1f50ea8r4brj21w01w0e84",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f50ea8r4brj21w01w0e84.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 4575232
                }
            }, {
                "pid": "4b7a8989jw1f50fvq2tcoj21a31dwnpf",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f50fvq2tcoj21a31dwnpf.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 3812352
                }
            }],
            "like_count": 42430,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-18 18:51",
            "id": 3987791681636863,
            "mid": "3987791681636863",
            "idstr": "3987791681636863",
            "text": "\u521a\u53d1\u73b0\u571f\u8c46\u62ff\u6211\u7684\u624b\u673a\u62cd\u4e86\u597d\u4e9b\u56fe\u7247\uff0c\u66f9\u8001\u8bf4\u50cf\u662f\u5c0f\u52a8\u7269\u7684\u89c6\u89d2\uff0c\u53d6\u666f\u89d2\u5ea6\u51fa\u4e4e\u610f\u6599\u3002\u6211\u548b\u89c9\u5f97\u753b\u9762\u8be1\u5f02\uff0c\u62cd\u5f97\u8ddf\u72af\u7f6a\u73b0\u573a\u4f3c\u7684\u2026<i class=\"face face_1 icon_20\">[doge]<\/i><i class=\"face face_4 icon_12\">[\u7167\u76f8\u673a]<\/i>",
            "textLength": 128,
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": ["4b7a8989jw1f4zkk9mpbzj22hm340e81", "4b7a8989jw1f4zkk7yco8j22hm340hdt", "4b7a8989jw1f4zkkazf6bj22hm340qv5", "4b7a8989jw1f4zkkcsgs3j22hm340u0x", "4b7a8989jw1f4zkkeeg6dj22hm340e81", "4b7a8989jw1f4zkkgx4qzj22hm340kjm"],
            "thumbnail_pic": "http:\/\/ww1.sinaimg.cn\/thumbnail\/4b7a8989jw1f4zkk9mpbzj22hm340e81.jpg",
            "bmiddle_pic": "http:\/\/ww1.sinaimg.cn\/bmiddle\/4b7a8989jw1f4zkk9mpbzj22hm340e81.jpg",
            "original_pic": "http:\/\/ww1.sinaimg.cn\/large\/4b7a8989jw1f4zkk9mpbzj22hm340e81.jpg",
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "picStatus": "0:1,1:1,2:1,3:1,4:1,5:1",
            "reposts_count": 428,
            "comments_count": 1653,
            "attitudes_count": 51966,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 4294967300,
            "hot_weibo_tags": [{
                "tag_name": "\u660e\u661f",
                "tag_scheme": "sinaweibo:\/\/cardlist?containerid=102803_ctg1_4288_-_ctg1_4288&extparam=from_mixbottomtag_-_tagmid_3987791681636863",
                "tag_hidden": 1,
                "tag_type": 2,
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2015\/12\/02\/142\/timeline_icon_hot.png",
                "tag_weight": 0.5,
                "from_cateid": "4288",
                "containerid": "102803_ctg1_4288_-_ctg1_4288"
            }],
            "text_tag_tips": [],
            "userType": 0,
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "tag_struct": [{
                "tag_name": "\u660e\u661f",
                "tag_scheme": "sinaweibo:\/\/cardlist?containerid=102803_ctg1_4288_-_ctg1_4288&extparam=from_mixbottomtag_-_tagmid_3987791681636863",
                "tag_hidden": 1,
                "tag_type": 2,
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2016\/03\/15\/196\/timeline_icon_hot.png",
                "oid": "1022:102803_ctg1_4288_-_ctg1_4288"
            }],
            "created_timestamp": 1466247089,
            "bid": "DAPRu6RP1",
            "pics": [{
                "pid": "4b7a8989jw1f4zkk9mpbzj22hm340e81",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f4zkk9mpbzj22hm340e81.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1429504
                }
            }, {
                "pid": "4b7a8989jw1f4zkk7yco8j22hm340hdt",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f4zkk7yco8j22hm340hdt.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1524736
                }
            }, {
                "pid": "4b7a8989jw1f4zkkazf6bj22hm340qv5",
                "url": "http:\/\/ww4.sinaimg.cn\/thumb180\/4b7a8989jw1f4zkkazf6bj22hm340qv5.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1810432
                }
            }, {
                "pid": "4b7a8989jw1f4zkkcsgs3j22hm340u0x",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f4zkkcsgs3j22hm340u0x.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1905664
                }
            }, {
                "pid": "4b7a8989jw1f4zkkeeg6dj22hm340e81",
                "url": "http:\/\/ww3.sinaimg.cn\/thumb180\/4b7a8989jw1f4zkkeeg6dj22hm340e81.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1429504
                }
            }, {
                "pid": "4b7a8989jw1f4zkkgx4qzj22hm340kjm",
                "url": "http:\/\/ww3.sinaimg.cn\/thumb180\/4b7a8989jw1f4zkkgx4qzj22hm340kjm.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 2668544
                }
            }],
            "like_count": 51966,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-18 01:41",
            "id": 3987532569810163,
            "mid": "3987532569810163",
            "idstr": "3987532569810163",
            "text": "\u4eca\u513f\u5728\u673a\u573a\u5475\u65a5\u4e86\u571f\u8c46\uff0c\u8fd8\u72e0\u72e0\u6253\u4e86\u4ed6\u5c41\u80a1\u4e00\u4e0b\u3002\u5c24\u5176\u662f\u540e\u6765\u624d\u5f97\u77e5\uff0c\u5176\u5b9e\u662f\u524d\u9762\u7684\u4eba\u63d2\u961f\uff0c\u624d\u95f4\u63a5\u5bfc\u81f4\u571f\u8c46\u5411\u524d\u63a8\u8f66\u649e\u5230\u4e86\u5bf9\u65b9\u7684\u63a8\u8f66\uff0c\u5fc3\u91cc\u66f4\u662f\u8d1f\u759a\u96be\u8fc7\u5230\u4e86\u73b0\u5728\u3002\u5982\u4f55\u5728\u7d27\u5f20\u6df7\u4e71\u7684\u516c\u4f17\u573a\u5408\u4fdd\u6301\u4e0d\u7126\u8e81\uff0c\u6b63\u786e\u5904\u7406\u5b69\u5b50\u7684\u987d\u76ae\u548c\u54ed\u95f9\uff0c\u8fd9\u95e8\u8bfe\u7a0b\u6211\u5f97\u4e86\u8d1f\u5206\u2026<i class=\"face face_2 icon_26\">[\u732a\u5934]<\/i>",
            "textLength": 236,
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": ["4b7a8989jw1f4yql6q3f3j22c02c0qv5"],
            "thumbnail_pic": "http:\/\/ww1.sinaimg.cn\/thumbnail\/4b7a8989jw1f4yql6q3f3j22c02c0qv5.jpg",
            "bmiddle_pic": "http:\/\/ww1.sinaimg.cn\/bmiddle\/4b7a8989jw1f4yql6q3f3j22c02c0qv5.jpg",
            "original_pic": "http:\/\/ww1.sinaimg.cn\/large\/4b7a8989jw1f4yql6q3f3j22c02c0qv5.jpg",
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "picStatus": "0:1",
            "reposts_count": 897,
            "comments_count": 2925,
            "attitudes_count": 80381,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 4294967300,
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "userType": 0,
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "created_timestamp": 1466185312,
            "bid": "DAJ7yFa4r",
            "pics": [{
                "pid": "4b7a8989jw1f4yql6q3f3j22c02c0qv5",
                "url": "http:\/\/ww1.sinaimg.cn\/wap180\/4b7a8989jw1f4yql6q3f3j22c02c0qv5.jpg",
                "size": "wap180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1810432
                }
            }],
            "like_count": 80381,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-16 22:16",
            "id": 3987118580279236,
            "mid": "3987118580279236",
            "idstr": "3987118580279236",
            "text": "\u81ea\u62cd\u5408\u5f71\u65f6\uff0c\u8138\u5927\u7684\u5bb9\u6613\u51fa\u753b\uff0c\u8138\u592a\u5c0f\u5bb9\u6613\u88ab\u6324\u7740\uff0c\u5634\u5927\u548b\u7b11\u90fd\u5f88\u660e\u663e\u2026\u2026<i class=\"face face_1 icon_20\">[doge]<\/i><i class=\"face face_4 icon_12\">[\u7167\u76f8\u673a]<\/i>",
            "textLength": 80,
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": ["4b7a8989gw1f4xfbbqxfcj20xc0qoh6f", "4b7a8989gw1f4xfbg9rp4j20xc0qoe0s", "4b7a8989gw1f4xfbmrjn2j20xc0qoe06"],
            "thumbnail_pic": "http:\/\/ww4.sinaimg.cn\/thumbnail\/4b7a8989gw1f4xfbbqxfcj20xc0qoh6f.jpg",
            "bmiddle_pic": "http:\/\/ww4.sinaimg.cn\/bmiddle\/4b7a8989gw1f4xfbbqxfcj20xc0qoh6f.jpg",
            "original_pic": "http:\/\/ww4.sinaimg.cn\/large\/4b7a8989gw1f4xfbbqxfcj20xc0qoh6f.jpg",
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "picStatus": "0:1,1:1,2:1",
            "reposts_count": 1948,
            "comments_count": 2855,
            "attitudes_count": 85011,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 4294967300,
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "userType": 0,
            "extend_info": {
                "weibo_camera": {
                    "c": ["29625552_29625553_29625554"]
                }
            },
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "created_timestamp": 1466086609,
            "bid": "DAylQ1aDO",
            "pics": [{
                "pid": "4b7a8989gw1f4xfbbqxfcj20xc0qoh6f",
                "url": "http:\/\/ww4.sinaimg.cn\/thumb180\/4b7a8989gw1f4xfbbqxfcj20xc0qoh6f.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 777681
                }
            }, {
                "pid": "4b7a8989gw1f4xfbg9rp4j20xc0qoe0s",
                "url": "http:\/\/ww3.sinaimg.cn\/thumb180\/4b7a8989gw1f4xfbg9rp4j20xc0qoe0s.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 782708
                }
            }, {
                "pid": "4b7a8989gw1f4xfbmrjn2j20xc0qoe06",
                "url": "http:\/\/ww3.sinaimg.cn\/thumb180\/4b7a8989gw1f4xfbmrjn2j20xc0qoe06.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 760180
                }
            }],
            "like_count": 85011,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-16 10:46",
            "id": 3986944842900373,
            "mid": "3986944842900373",
            "idstr": "3986944842900373",
            "text": "\u770b\u5230\u8fd9\u7ec4\u56fe\u7247\uff0c\u6606\u4ed1\u96ea\u5c71\uff0c\u7cbe\u7edd\u53e4\u57ce\uff0c\u6614\u65e5\u60c5\u666f\u518d\u73b0\uff0c\u7ecf\u5386\u8fc7\u7684\u5947\u5e7b\u5386\u5386\u5728\u76ee\u3002<a class='k' href='\/k\/\u9b3c\u5439\u706f3D?from=feed'>#\u9b3c\u5439\u706f3D#<\/a>\u624b\u6e38\u4eca\u65e5\u5168\u5e73\u53f0\u4e0a\u7ebf\uff0cshirley\u6768\u7684\u5192\u9669\u8fd8\u5c06\u7ee7\u7eed\uff01\u522b\u7b49\u5f85\uff0c\u8ddf\u6211\u6765\uff01\u6233<a data-url=\"http:\/\/t.cn\/R5alISA\" href=\"http:\/\/weibo.cn\/sinaurl?u=http%!A(MISSING)%!F(MISSING)%!F(MISSING)t.cn%!F(MISSING)R5alISA\" ><i class=\"iconimg iconimg-xs\"><img src=\"http:\/\/h5.sinaimg.cn\/upload\/2015\/09\/25\/3\/timeline_card_small_web_default.png\"><\/i><span class=\"surl-text\">\u7f51\u9875\u94fe\u63a5<\/span><\/a>\uff0c\u4e00\u8d77\u8fdb\u5165\u8fd9\u4e1c\u65b9\u7075\u5f02\u63a2\u9669\u624b\u6e38\u7684\u9b54\u5e7b\u4e16\u754c\uff01[\u5f20\u5634]<i class=\"face face_3 icon_9\">[haha]<\/i><a href='\/n\/\u9b3c\u5439\u706f3D'>@\u9b3c\u5439\u706f3D<\/a>",
            "textLength": 221,
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": ["4b7a8989jw1f4wv5y2c3uj20hs0vk77u", "4b7a8989jw1f4wv5y9xzfj20dc0m8wg6", "4b7a8989jw1f4wv5xuwquj20k00zkjvx", "4b7a8989jw1f4wv5ztq3kj20rs119e81"],
            "thumbnail_pic": "http:\/\/ww2.sinaimg.cn\/thumbnail\/4b7a8989jw1f4wv5y2c3uj20hs0vk77u.jpg",
            "bmiddle_pic": "http:\/\/ww2.sinaimg.cn\/bmiddle\/4b7a8989jw1f4wv5y2c3uj20hs0vk77u.jpg",
            "original_pic": "http:\/\/ww2.sinaimg.cn\/large\/4b7a8989jw1f4wv5y2c3uj20hs0vk77u.jpg",
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "picStatus": "0:1,1:1,2:1,3:1",
            "reposts_count": 3108,
            "comments_count": 3469,
            "attitudes_count": 40328,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 4294967300,
            "expire_time": 1466144457,
            "page_type": 32,
            "hot_weibo_tags": [{
                "tag_name": "\u6e38\u620f",
                "tag_scheme": "sinaweibo:\/\/cardlist?containerid=102803_ctg1_4888_-_ctg1_4888&extparam=from_mixbottomtag_-_tagmid_3986944842900373",
                "tag_hidden": 1,
                "tag_type": 2,
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2015\/12\/02\/142\/timeline_icon_hot.png",
                "tag_weight": 0.69999694824219,
                "from_cateid": "4888",
                "containerid": "102803_ctg1_4888_-_ctg1_4888"
            }, {
                "tag_name": "\u624b\u673a\u6e38\u620f",
                "tag_scheme": "sinaweibo:\/\/cardlist?containerid=102803_ctg1_200263_-_ctg1_200263&extparam=from_mixbottomtag_-_tagmid_3986944842900373",
                "tag_hidden": 1,
                "tag_type": 2,
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2015\/12\/02\/142\/timeline_icon_hot.png",
                "tag_weight": 0.64999389648438,
                "from_cateid": "200263",
                "containerid": "102803_ctg1_200263_-_ctg1_200263"
            }, {
                "tag_name": "\u63a2\u9669",
                "tag_scheme": "sinaweibo:\/\/cardlist?containerid=102803_ctg1_a1c57eed0fcef666b6a02c1dc9284ab9_-_ctg1_a1c57eed0fcef666b6a02c1dc9284ab9&extparam=from_mixbottomtag_-_tagmid_3986944842900373",
                "tag_hidden": 1,
                "tag_type": 2,
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2015\/12\/02\/142\/timeline_icon_hot.png",
                "tag_weight": 0.50099182128906,
                "from_cateid": "a1c57eed0fcef666b6a02c1dc9284ab9",
                "containerid": "102803_ctg1_a1c57eed0fcef666b6a02c1dc9284ab9_-_ctg1_a1c57eed0fcef666b6a02c1dc9284ab9"
            }],
            "text_tag_tips": [],
            "userType": 0,
            "extend_info": {
                "weibo_camera": {
                    "c": ["21564352_26220801_19760418_28282499"]
                },
                "ad": {
                    "url_marked": "true"
                }
            },
            "positive_recom_flag": 0,
            "mblogtype": 1,
            "mark": "followtopweibo",
            "url_struct": [{
                "short_url": "http:\/\/t.cn\/R5alISA",
                "ori_url": "http:\/\/t.cn\/R5alISA",
                "url_title": "\u7f51\u9875\u94fe\u63a5",
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2015\/09\/25\/3\/timeline_card_small_web.png",
                "position": 2,
                "url_type": 0,
                "result": false,
                "need_save_obj": 1
            }],
            "topic_struct": [{
                "topic_title": "\u9b3c\u5439\u706f3D",
                "topic_url": "sinaweibo:\/\/pageinfo?containerid=10080838906c1580b1ae9e3cda0b4bb8cd30ec&pageid=10080838906c1580b1ae9e3cda0b4bb8cd30ec&extparam=%!E(MISSING)9%!A(MISSING)C%!B(MISSING)C%!E(MISSING)5%B9%!E(MISSING)7%AF3D"
            }],
            "page_info": {
                "page_id": "10080838906c1580b1ae9e3cda0b4bb8cd30ec",
                "type": 0,
                "page_title": "#\u9b3c\u5439\u706f3D#",
                "page_url": "http:\/\/m.weibo.cn\/p\/index?containerid=10080838906c1580b1ae9e3cda0b4bb8cd30ec",
                "page_pic": "http:\/\/ww1.sinaimg.cn\/thumb300\/eab7d77bgw1f4twkhd8rzj2050050dgg.jpg",
                "page_desc": "\u4e09\u4ebf\u706f\u4e1d\u671f\u5f85\uff0c\u59da\u6668\u503e\u60c5\u4ee3\u8a00\uff01\u7531\u65b0\u52a8\u96c6\u56e2\u72ec\u5bb6\u4ee3\u7406\u300165\u6e38\u620f\u8054\u5408\u53d1\u884c\u7684\u300a\u9b3c\u5439\u706f3D\u300b\u6b63\u7248\u6388\u6743\u624b\u6e38\uff0c\u5c06\u4e3a\u706f\u4e1d\u548c\u73a9\u5bb6\u91cd\u65b0\u5f00\u542f\u4e00\u4e2a\u7075\u5f02\u4e0e\u8650\u5fc3\u7684\u9b3c\u9b45\u4e16\u754c\uff0c\u8d85\u4eba\u6c14\u660e\u661f\u59da\u6668\u52a0\u76df\u300a\u9b3c\u5439\u706f3D\u300b\u624b\u6e38\uff0c\u5e76\u5206\u9970\u5973\u4e3b\u89d2Shirley\u6768\u548c\u7ec8\u6781BOSS\u7cbe\u7edd\u5973\u738b\u3002\u6781\u81f4\u8fd8\u539f\uff0c\u60ca\u609a\u8650\u5fc3\uff01",
                "object_type": "topic",
                "tips": "21\u4eba\u5173\u6ce8",
                "object_id": "1022:10080838906c1580b1ae9e3cda0b4bb8cd30ec",
                "actionlog": {
                    "oid": "sinaweibo:\/\/pageinfo?containerid=10080838906c1580b1ae9e3cda0b4bb8cd30ec&containerid=10080838906c1580b1ae9e3cda0b4bb8cd30ec&extparam=%!E(MISSING)9%!A(MISSING)C%!B(MISSING)C%!E(MISSING)5%B9%!E(MISSING)7%AF3D",
                    "act_code": 300,
                    "act_type": 1,
                    "ext": "mid:3986944842900373|rid:",
                    "mid": "3986944842900373",
                    "source": "ad",
                    "code": "14000014",
                    "mark": "followtopweibo"
                },
                "content1": "",
                "content2": ""
            },
            "tag_struct": [{
                "tag_name": "\u6e38\u620f",
                "tag_scheme": "sinaweibo:\/\/cardlist?containerid=102803_ctg1_4888_-_ctg1_4888&extparam=from_mixbottomtag_-_tagmid_3986944842900373",
                "tag_hidden": 1,
                "tag_type": 2,
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2016\/03\/15\/196\/timeline_icon_hot.png",
                "oid": "1022:102803_ctg1_4888_-_ctg1_4888"
            }, {
                "tag_name": "\u624b\u673a\u6e38\u620f",
                "tag_scheme": "sinaweibo:\/\/cardlist?containerid=102803_ctg1_200263_-_ctg1_200263&extparam=from_mixbottomtag_-_tagmid_3986944842900373",
                "tag_hidden": 1,
                "tag_type": 2,
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2016\/03\/15\/196\/timeline_icon_hot.png",
                "oid": "1022:102803_ctg1_200263_-_ctg1_200263"
            }, {
                "tag_name": "\u63a2\u9669",
                "tag_scheme": "sinaweibo:\/\/cardlist?containerid=102803_ctg1_a1c57eed0fcef666b6a02c1dc9284ab9_-_ctg1_a1c57eed0fcef666b6a02c1dc9284ab9&extparam=from_mixbottomtag_-_tagmid_3986944842900373",
                "tag_hidden": 1,
                "tag_type": 2,
                "url_type_pic": "http:\/\/h5.sinaimg.cn\/upload\/2016\/03\/15\/196\/timeline_icon_hot.png",
                "oid": "1022:102803_ctg1_a1c57eed0fcef666b6a02c1dc9284ab9_-_ctg1_a1c57eed0fcef666b6a02c1dc9284ab9"
            }],
            "created_timestamp": 1466045187,
            "bid": "DAtPCcawd",
            "pics": [{
                "pid": "4b7a8989jw1f4wv5y2c3uj20hs0vk77u",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f4wv5y2c3uj20hs0vk77u.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 141498
                }
            }, {
                "pid": "4b7a8989jw1f4wv5y9xzfj20dc0m8wg6",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f4wv5y9xzfj20dc0m8wg6.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 72610
                }
            }, {
                "pid": "4b7a8989jw1f4wv5xuwquj20k00zkjvx",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f4wv5xuwquj20k00zkjvx.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 177710
                }
            }, {
                "pid": "4b7a8989jw1f4wv5ztq3kj20rs119e81",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f4wv5ztq3kj20rs119e81.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1429504
                }
            }],
            "like_count": 40328,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-15 11:21",
            "id": 3986591343160716,
            "mid": "3986591343160716",
            "idstr": "3986591343160716",
            "text": "\u505a\u86cb\u7cd5\uff0c\u901b\u8d85\u5e02\uff0c\u4f30\u8ba1\u4e0d\u4e45\u571f\u8c46\u5c31\u53ef\u4ee5\u72ec\u7acb\u751f\u6d3b\u4e86\u3002<i class=\"face face_1 icon_20\">[doge]<\/i>[\u5f20\u5634]",
            "textLength": 56,
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": ["4b7a8989jw1f4vqgdate9j21w02io4qq", "4b7a8989jw1f4vqq49vlwj21w02ioe82", "4b7a8989jw1f4vqhcxp3wj22c02c0e81"],
            "thumbnail_pic": "http:\/\/ww3.sinaimg.cn\/thumbnail\/4b7a8989jw1f4vqgdate9j21w02io4qq.jpg",
            "bmiddle_pic": "http:\/\/ww3.sinaimg.cn\/bmiddle\/4b7a8989jw1f4vqgdate9j21w02io4qq.jpg",
            "original_pic": "http:\/\/ww3.sinaimg.cn\/large\/4b7a8989jw1f4vqgdate9j21w02io4qq.jpg",
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "picStatus": "0:1,1:1,2:1",
            "reposts_count": 225,
            "comments_count": 1737,
            "attitudes_count": 41310,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 4294967300,
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "userType": 0,
            "extend_info": {
                "weibo_camera": {
                    "c": ["32087104_25494305"]
                }
            },
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "created_timestamp": 1465960906,
            "bid": "DAkDsdgfi",
            "pics": [{
                "pid": "4b7a8989jw1f4vqgdate9j21w02io4qq",
                "url": "http:\/\/ww3.sinaimg.cn\/thumb180\/4b7a8989jw1f4vqgdate9j21w02io4qq.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 2192384
                }
            }, {
                "pid": "4b7a8989jw1f4vqq49vlwj21w02ioe82",
                "url": "http:\/\/ww1.sinaimg.cn\/thumb180\/4b7a8989jw1f4vqq49vlwj21w02ioe82.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 2478080
                }
            }, {
                "pid": "4b7a8989jw1f4vqhcxp3wj22c02c0e81",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989jw1f4vqhcxp3wj22c02c0e81.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1429504
                }
            }],
            "like_count": 41310,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-15 10:19",
            "id": 3986575614884440,
            "mid": "3986575614884440",
            "idstr": "3986575614884440",
            "text": "<i class=\"face face_3 icon_12\">[\u9c9c\u82b1]<\/i> \/\/<a href='\/n\/\u7535\u5f71\u4eba\u7a0b\u9752\u677e'>@\u7535\u5f71\u4eba\u7a0b\u9752\u677e<\/a>:\u4eca\u5929\u662f\u6211\u519c\u5386\u751f\u65e5\uff0c\u5373\u5c06\u63a8\u51fa\u7b2c\u4e03\u8f91\u300a\u5f71\u53f2100\u5927\u9752\u6625\u7247\u300b\uff0c\u656c\u8bf7\u5173\u6ce8\u3002<a href='\/n\/\u59da\u6668'>@\u59da\u6668<\/a> ",
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": [],
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "pid": 3986420156763881,
            "retweeted_status": {
                "created_at": "2015-07-10 01:17",
                "id": 3862865008050751,
                "mid": "3862865008050751",
                "idstr": "3862865008050751",
                "text": "\u63a8\u8350\u5173\u6ce8\uff1a\u9752\u5e74\u7535\u5f71\u624b\u518c\u5fae\u4fe1\u516c\u4f17\u53f7 dianyingshouce66",
                "source_allowclick": 0,
                "source_type": 1,
                "source": "Smartisan T1",
                "favorited": false,
                "pic_ids": [],
                "user": {
                    "id": 1662717013,
                    "screen_name": "\u8d3e\u6a1f\u67ef",
                    "profile_image_url": "http:\/\/tva4.sinaimg.cn\/crop.256.40.768.768.180\/631b0c55jw1etg9dwf5nwj20zk0nowie.jpg",
                    "profile_url": "\/u\/1662717013",
                    "statuses_count": 5911,
                    "verified": true,
                    "verified_reason": "\u5bfc\u6f14",
                    "description": "",
                    "remark": "",
                    "verified_type": 0,
                    "gender": "m",
                    "mbtype": 12,
                    "h5icon": {
                        "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                        "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                    },
                    "ismember": 1,
                    "valid": null,
                    "fansNum": "1616\u4e07",
                    "follow_me": false,
                    "following": false
                },
                "reposts_count": 8315,
                "comments_count": 571,
                "attitudes_count": 1225,
                "isLongText": false,
                "mlevel": 0,
                "visible": {
                    "type": 0,
                    "list_id": 0
                },
                "biz_feature": 0,
                "hot_weibo_tags": [],
                "text_tag_tips": [],
                "userType": 0,
                "positive_recom_flag": 0,
                "created_timestamp": 1436462249,
                "bid": "CqnpaxMmP",
                "like_count": 1225,
                "attitudes_status": 0
            },
            "reposts_count": 37,
            "comments_count": 343,
            "attitudes_count": 7812,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 0,
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "userType": 0,
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "created_timestamp": 1465957155,
            "bid": "DAke5kuFi",
            "like_count": 7812,
            "attitudes_status": 0
        }
    }, {
        "card_type": 9,
        "mblog": {
            "created_at": "06-14 22:48",
            "id": 3986401861506364,
            "mid": "3986401861506364",
            "idstr": "3986401861506364",
            "text": "\u5165\u591c\u7684\u57ce\uff0c\u88ab\u661f\u70b9\u7684\u5149\u71c3\u70e7\u7740\u6709\u4e86\u6e29\u5ea6\u3002\ud83c\udf1b",
            "textLength": 40,
            "source_allowclick": 0,
            "source_type": 1,
            "source": "iPhone 6s",
            "favorited": false,
            "pic_ids": ["4b7a8989gw1f4v51gxfinj22c02c0b29", "4b7a8989gw1f4v50ffq7qj23402c0b2a", "4b7a8989gw1f4v510z934j23402c0e82"],
            "thumbnail_pic": "http:\/\/ww4.sinaimg.cn\/thumbnail\/4b7a8989gw1f4v51gxfinj22c02c0b29.jpg",
            "bmiddle_pic": "http:\/\/ww4.sinaimg.cn\/bmiddle\/4b7a8989gw1f4v51gxfinj22c02c0b29.jpg",
            "original_pic": "http:\/\/ww4.sinaimg.cn\/large\/4b7a8989gw1f4v51gxfinj22c02c0b29.jpg",
            "user": {
                "id": 1266321801,
                "screen_name": "\u59da\u6668",
                "profile_image_url": "http:\/\/tva3.sinaimg.cn\/crop.0.1.509.509.180\/4b7a8989jw8est6r45x8ij20e50e8mxz.jpg",
                "profile_url": "\/u\/1266321801",
                "statuses_count": 8774,
                "verified": true,
                "verified_reason": "\u6f14\u5458\uff0c\u8054\u5408\u56fd\u96be\u6c11\u7f72\u4e2d\u56fd\u4eb2\u5584\u5927\u4f7f\u3002",
                "description": "\u5de5\u4f5c\u4e8b\u5b9c\u8bf7\u8054\u7edc\u7ecf\u7eaa\u4eba\u5f20\u857e\uff1arunrunlei@126.com \ud83d\ude0a",
                "remark": "",
                "verified_type": 0,
                "gender": "f",
                "mbtype": 12,
                "h5icon": {
                    "main": "http:\/\/u1.sinaimg.cn\/upload\/2013\/02\/22\/v_yellow_2x.png",
                    "other": ["http:\/\/u1.sinaimg.cn\/upload\/2013\/01\/23\/crown_2x.png"]
                },
                "ismember": 1,
                "valid": null,
                "fansNum": "7950\u4e07",
                "follow_me": false,
                "following": true
            },
            "picStatus": "0:1,1:1,2:1",
            "reposts_count": 366,
            "comments_count": 1343,
            "attitudes_count": 37981,
            "isLongText": false,
            "mlevel": 0,
            "visible": {
                "type": 0,
                "list_id": 0
            },
            "biz_feature": 4294967300,
            "hot_weibo_tags": [],
            "text_tag_tips": [],
            "userType": 0,
            "extend_info": {
                "weibo_camera": {
                    "c": ["28676017_28676018"]
                }
            },
            "positive_recom_flag": 0,
            "mblogtype": 0,
            "created_timestamp": 1465915730,
            "bid": "DAfHQ6jSc",
            "pics": [{
                "pid": "4b7a8989gw1f4v51gxfinj22c02c0b29",
                "url": "http:\/\/ww4.sinaimg.cn\/thumb180\/4b7a8989gw1f4v51gxfinj22c02c0b29.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 1334272
                }
            }, {
                "pid": "4b7a8989gw1f4v50ffq7qj23402c0b2a",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989gw1f4v50ffq7qj23402c0b2a.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 2382848
                }
            }, {
                "pid": "4b7a8989gw1f4v510z934j23402c0e82",
                "url": "http:\/\/ww2.sinaimg.cn\/thumb180\/4b7a8989gw1f4v510z934j23402c0e82.jpg",
                "size": "thumb180",
                "geo": {
                    "width": 180,
                    "height": 180,
                    "croped": false,
                    "byte": 2478080
                }
            }],
            "like_count": 37981,
            "attitudes_status": 0
        }
    }]
}`
	var data result.WBRawPosts
	err := json.Unmarshal([]byte(src), &data)
	if err != nil {
		log.Println("err: ", err)
	}
	body, _ := json.Marshal(data)
	log.Printf("%+v", string(body))

}
func Decode() {
	url := "http://mp.weixin.qq.com/s?src=3&amp;timestamp=1466051965&amp;ver=1&amp;signature=1YSeL*0wzSX-IYC-oABCW6*rKniaRYV7Zh7FJ1lQnl5BAFJwaJD7Jptz5PgLwCFMbZGVIX-ajeuTRbou*h4umb4QGEj0ImG6SsWjqdH1g2EST9nrHNqUKD9U3UBWhh1vfkjydHYqvhgHSeRns-oWroFmmfq76oxOiFKH1DrniIw="
	fmt.Printf("%s", url)
}
func TimeParse() {
	t, err := time.Parse("2006-01-02T15:04:05", "2016-02-01T12:01:03")
	if err != nil {
		log.Println(err)
	}
	log.Println(t)
}
func MongoQuery() {
	session, err := mgo.Dial("192.168.20.24:27000")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("Book")
	//ids := []int{3119, 3072, 3070, 3060}
	//for _, i := range ids {
	for i := 7106; i > 0; i-- {

		url := "http://it-ebooks.info/book/" + strconv.Itoa(i) + "/"
		count, err := c.Find(bson.M{"link": url}).Count()
		if err != nil {
			log.Fatal(err)
		}
		if count != 1 {
			log.Fatal(url)
		} else {
			log.Println(i)
		}

		//time.Sleep(200 * time.Millisecond)
		//	http://it-ebooks.info/book/3119/
		//	http://it-ebooks.info/book/3072/
		//	http://it-ebooks.info/book/3070/
		//http://it-ebooks.info/book/3060/
	}

	//fmt.Println("Phone:", result.)
}
func RequestDemo() {
	agent := gorequest.New()
	go func() {
		for {
			response, _, _ := agent.Get("http://localhost:10086/api/wb/profile/1732447702").End()
			log.Println(response.Body)
		}
	}()

	handleSignals()
}
func handleSignals() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
func URLTest() {
	rawurl := "http:/www.facebook.com/"
	realUrl, err := url.Parse(rawurl)
	if err != nil {
		log.Println("error: ", err)
	} else {
		log.Println("real url : ", realUrl.RequestURI())
	}
}

func WaitGroupDemo() {
	messages := make(chan int)
	var wg sync.WaitGroup

	// you can also add these one at
	// a time if you need to

	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)
		messages <- 5
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		messages <- 3
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 7)
		messages <- 7
	}()
	go func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()
	wg.Wait()

}

func GoRoutine() {
	respon := make(chan string)
	go func() {
		respon <- "test1"
		time.Sleep(time.Second)
	}()

	go func() {
		respon <- "test2"
		time.Sleep(time.Second)
	}()
	go func() {
		respon <- "test2"
		time.Sleep(time.Second)
	}()
	go func() {
		respon <- "test1"
		time.Sleep(time.Second)
	}()
	for {
		select {
		case t := <-respon:
			log.Println(t)
		}

		time.Sleep(time.Second)
	}
}

func StartTokenGen() {
	tokenCh = make(chan string)
	popchan = make(chan string)

	go TokenGen(tokenCh)
	//go TokenGen1(popchan)

	for {
		select {
		case t := <-tokenCh:
			log.Println(t)
		case p := <-popchan:
			log.Println(p)
		default:
			log.Println("no message")
		}
	}

}

func TokenGen1(ch chan string) {
	for {
		popchan <- "pop"
		time.Sleep(1 * time.Second)

	}
}

func TokenGen(ch chan string) {
	for {
		log.Println("goroutine")
		time.Sleep(1 * time.Second)

	}
}
func RegexpDemo() {
	r, _ := regexp.Compile(`微信号: (\S+)<`)

	matcher := r.FindStringSubmatch("12412微信号: Bit_baike1241231")

	fmt.Println(matcher)
}
