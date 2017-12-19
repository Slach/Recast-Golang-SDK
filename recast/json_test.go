package recast

func getBadRequestJSONResponse() string {
	return `{"results":null,"message":"Request is invalid"}`
}

func getServerErrorJSONResponse() string {
	return `{"results":null,"message":"Internal server error"}`
}

func getSuccessfulPostMessageResponse() string {
	return `{"results":null,"message":"Messages successfully posted"}`
}

func getSuccessfulConverseJSONResponse() string {
	return `{
		"results":{
			"uuid": "390ff8ee-2909-4da4-bdae-15914ad2eacb",
			"source": "projet",
			"replies": [
			"Quel est ton projet cette année ?"
			],
			"action": {
				"slug": "projet",
				"done": false,
				"reply": "Quel est ton projet cette année ?"
			},
			"next_actions": [],
			"memory": {
				"secteur-boite": null,
				"password": null,
				"email": null,
				"confirmation": null,
				"context": null,
				"moyen-contact": null,
				"domaines-aides": null,
				"needs": null,
				"job": null,
				"type_projet": null,
				"numero-etudiant": null
			},
			"entities": {},
			"intents": [{
				"confidence": 0.99,
				"slug": "projet"
			}],
			"conversation_token": "8480ed074ab739198e709e46ef226420",
			"language": "fr",
			"timestamp": "2017-03-23T14:11:16.111655+00:00",
			"version": "2.4.0",
			"status": 200
		}`
}

func getBadFormatJSONResponse() string {
	return `{
   "results":{
	  "uuid": "7c88d59d-9eaa-4b4f-ba3d-be466cf03b5f",
      "source":"WhatistheweatherinLondontomorrow?AndinParis?",
      "intents":[
         {
            "slug":"weather",
            "confidence":0.67
         }
      ,
    }`
}

func getSuccessfulRequestJSONResponse() string {
	return `{
   "results":{
	   "uuid": "7c88d59d-9eaa-4b4f-ba3d-be466cf03b5f",
      "source":"WhatistheweatherinLondontomorrow?AndinParis?",
      "intents":[
         {
            "slug":"weather",
            "confidence":0.67
         }
      ],
      "act":"wh-query",
      "type":"desc:desc",
      "sentiment":"neutral",
      "entities":{
         "action":[
            {
               "agent":"theweatherinLondon",
               "tense":"present",
               "raw":"is",
               "confidence":0.89
            }
         ],
         "location":[
            {
               "formated":"London,London,GreaterLondon,England,UnitedKingdom",
               "lat":51.5073509,
               "lng":-0.1277583,
               "raw":"London",
               "confidence":0.97
            },
            {
               "formated":"Paris,Paris,Île-de-France,France",
               "lat":48.856614,
               "lng":2.3522219,
               "raw":"Paris",
               "confidence":0.83
            }
         ],
         "datetime":[
            {
               "value":"2016-07-11T10:00:00+00:00",
               "raw":"tomorrow",
               "confidence":0.95
            }
         ]
      },
	  "language":"en",
	  "version":"2.0.0",
	  "timestamp":"2016-10-12T15:34:57.298559Z",
	  "status":200
      }
    }`
}

func getSuccessNoIntentRequestJSONResponse() string {
	return `{
		"results": {
			"uuid": "7c88d59d-9eaa-4b4f-ba3d-be466cf03b5f",
			"source": "Some text",
			"intents": [],
			"act": "assert",
			"type": "desc:desc",
			"sentiment": "neutral",
			"entities": {},
			"language": "en",
			"version":"2.0.0",
			"timestamp":"2016-07-10T23:17:59+02:00",
			"status":200
		}
	}`
}

func getValidConversationMessage() string {
	return `
{
	"message": {
		"data": {
			"userName": "Yannick Heinrich"
		},
		"__v": 0,
		"participant": "c9244b31-45f2-431c-be10-d3361851cf7e",
		"conversation": "f206b482-cb0c-435b-91bc-4628c8829d83",
		"attachment": {
			"content": "Hello",
			"type": "text"
		},
		"receivedAt": "2017-03-20T21:58:52.346Z",
		"isActive": true,
		"_id": "61a7921b-f771-4211-82ca-05885160fd6d"
	},
	"chatId": "chatId",
	"senderId": "senderId"
}`
}

func getSuccessfulDialogJSONResponse() string {
	return `{
	"message": "OK",
	"results": {
	  "messages": [
		{
		  "type": "quickReplies",
		  "content": {
			"title": "Quick reply title",
			"buttons": [
			  {
				"value": "quick reply payload1",
				"title": "quick reply text1"
			  },
			  {
				"value": "quick reply payload2",
				"title": "quick reply text2"
			  }
			]
		  },
		  "value": {
			"title": "Quick reply title",
			"buttons": [
			  {
				"value": "quick reply payload1",
				"title": "quick reply text1"
			  },
			  {
				"value": "quick reply payload2",
				"title": "quick reply text2"
			  }
			]
		  }
		},
		{
		  "type": "text",
		  "content": "text message"
		},
		{
		  "type": "card",
		  "content": {
			"title": "card title",
			"subtitle": "card subtitle",
			"imageUrl": "https://beebom-redkapmedia.netdna-ssl.com/wp-content/uploads/2016/01/Reverse-Image-Search-Engines-Apps-And-Its-Uses-2016.jpg",
			"buttons": [
			  {
				"value": "http://google.com",
				"title": "link button",
				"type": "web_url"
			  },
			  {
				"value": "postback value",
				"title": "postback button",
				"type": "postback"
			  },
			  {
				"value": "+33101010101",
				"title": "phone button",
				"type": "phonenumber"
			  }
			]
		  }
		},
		{
		  "type": "card",
		  "content": {
			"title": "Button title",
			"buttons": [
			  {
				"value": "http://google.com",
				"title": "link button",
				"type": "web_url"
			  },
			  {
				"value": "postback payload",
				"title": "postback button",
				"type": "postback"
			  },
			  {
				"value": "+33101010101",
				"title": "phone button",
				"type": "phonenumber"
			  }
			]
		  }
		},
		{
		  "type": "carousel",
		  "content": [
			{
			  "title": "Card 1 title",
			  "subtitle": "Card 1 subtitle",
			  "imageUrl": "https://beebom-redkapmedia.netdna-ssl.com/wp-content/uploads/2016/01/Reverse-Image-Search-Engines-Apps-And-Its-Uses-2016.jpg",
			  "buttons": [
				{
				  "value": "postback payload",
				  "title": "postback button",
				  "type": "postback"
				}
			  ]
			},
			{
			  "title": "Card 2 title",
			  "subtitle": "Card 2 subtitle",
			  "imageUrl": "https://beebom-redkapmedia.netdna-ssl.com/wp-content/uploads/2016/01/Reverse-Image-Search-Engines-Apps-And-Its-Uses-2016.jpg",
			  "buttons": []
			}
		  ],
		  "value": [
			{
			  "title": "Card 1 title",
			  "subtitle": "Card 1 subtitle",
			  "imageUrl": "https://beebom-redkapmedia.netdna-ssl.com/wp-content/uploads/2016/01/Reverse-Image-Search-Engines-Apps-And-Its-Uses-2016.jpg",
			  "buttons": [
				{
				  "value": "postback payload",
				  "title": "postback button",
				  "type": "postback"
				}
			  ]
			},
			{
			  "title": "Card 2 title",
			  "subtitle": "Card 2 subtitle",
			  "imageUrl": "https://beebom-redkapmedia.netdna-ssl.com/wp-content/uploads/2016/01/Reverse-Image-Search-Engines-Apps-And-Its-Uses-2016.jpg",
			  "buttons": []
			}
		  ]
		},
		{
		  "type": "list",
		  "content": {
			"elements": [
			  {
				"title": "List 1 title",
				"imageUrl": "https://beebom-redkapmedia.netdna-ssl.com/wp-content/uploads/2016/01/Reverse-Image-Search-Engines-Apps-And-Its-Uses-2016.jpg",
				"subtitle": "List 1 subtitle",
				"buttons": [
				  {
					"title": "postback button",
					"value": "postback payload",
					"type": "postback"
				  }
				]
			  },
			  {
				"title": "List 2 title",
				"imageUrl": "https://beebom-redkapmedia.netdna-ssl.com/wp-content/uploads/2016/01/Reverse-Image-Search-Engines-Apps-And-Its-Uses-2016.jpg",
				"subtitle": "List 2 subtitle",
				"buttons": [
				  {
					"title": "link button",
					"value": "http://google.com",
					"type": "web_url"
				  }
				]
			  }
			],
			"buttons": []
		  }
		},
		{
		  "type": "picture",
		  "content": "https://beebom-redkapmedia.netdna-ssl.com/wp-content/uploads/2016/01/Reverse-Image-Search-Engines-Apps-And-Its-Uses-2016.jpg"
		}
	  ],
	  "conversation": {
		"id": "recast-test-1513612367047",
		"language": "en",
		"memory": {
		  "direction": {
			"bearing": 270,
			"raw": "west",
			"confidence": 0.58
		  }
		},
		"skill": "test",
		"skill_occurences": 2
	  },
	  "nlp": {
		"uuid": "d2782e9b-19a0-4d23-9326-7e6d2299a99f",
		"source": "west",
		"intents": [],
		"act": "assert",
		"type": null,
		"sentiment": "neutral",
		"entities": {
		  "cardinal": [
			{
			  "bearing": 270,
			  "raw": "west",
			  "confidence": 0.58
			}
		  ]
		},
		"language": "en",
		"processing_language": "en",
		"version": "2.11.0",
		"timestamp": "2017-12-18T15:53:04.949696+00:00",
		"status": 200
	  }
	}
}`
}
