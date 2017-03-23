package recast

func getBadRequestJSONResponse() string {
	return `{"results":null,"message":"Request is invalid"}`
}

func getServerErrorJSONResponse() string {
	return `{"results":null,"message":"Internal server error"}`
}

func getSuccessFulPostMessageResponse() {
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
