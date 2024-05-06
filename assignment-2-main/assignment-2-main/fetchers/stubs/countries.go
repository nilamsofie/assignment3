package stubs

func CountryByCode(code string) string {
	switch code {
	case "NO":
		return norwayData
	case "DE":
		return germanyData
	case "US":
		return usaData
	default:
		return emptyData
	}
}

func CountryByName(name string) string {
	switch name {
	case "Norway":
		return norwayData
	case "Germany":
		return germanyData
	case "USA":
		return usaData
	default:
		return emptyData
	}
}

const emptyData = `[]`

const norwayData = `
[
  {
    "name": {
      "common": "Norway",
      "official": "Kingdom of Norway",
      "nativeName": {
        "nno": {
          "official": "Kongeriket Noreg",
          "common": "Noreg"
        },
        "nob": {
          "official": "Kongeriket Norge",
          "common": "Norge"
        },
        "smi": {
          "official": "Norgga gonagasriika",
          "common": "Norgga"
        }
      }
    },
    "tld": [
      ".no"
    ],
    "cca2": "NO",
    "ccn3": "578",
    "cca3": "NOR",
    "cioc": "NOR",
    "independent": true,
    "status": "officially-assigned",
    "unMember": true,
    "currencies": {
      "NOK": {
        "name": "Norwegian krone",
        "symbol": "kr"
      }
    },
    "idd": {
      "root": "+4",
      "suffixes": [
        "7"
      ]
    },
    "capital": [
      "Oslo"
    ],
    "altSpellings": [
      "NO",
      "Norge",
      "Noreg",
      "Kingdom of Norway",
      "Kongeriket Norge",
      "Kongeriket Noreg"
    ],
    "region": "Europe",
    "subregion": "Northern Europe",
    "languages": {
      "nno": "Norwegian Nynorsk",
      "nob": "Norwegian Bokmål",
      "smi": "Sami"
    },
    "translations": {
      "ara": {
        "official": "مملكة النرويج",
        "common": "النرويج"
      },
      "bre": {
        "official": "Rouantelezh Norvegia",
        "common": "Norvegia"
      },
      "ces": {
        "official": "Norské království",
        "common": "Norsko"
      },
      "cym": {
        "official": "Kingdom of Norway",
        "common": "Norway"
      },
      "deu": {
        "official": "Königreich Norwegen",
        "common": "Norwegen"
      },
      "est": {
        "official": "Norra Kuningriik",
        "common": "Norra"
      },
      "fin": {
        "official": "Norjan kuningaskunta",
        "common": "Norja"
      },
      "fra": {
        "official": "Royaume de Norvège",
        "common": "Norvège"
      },
      "hrv": {
        "official": "Kraljevina Norveška",
        "common": "Norveška"
      },
      "hun": {
        "official": "Norvég Királyság",
        "common": "Norvégia"
      },
      "ita": {
        "official": "Regno di Norvegia",
        "common": "Norvegia"
      },
      "jpn": {
        "official": "ノルウェー王国",
        "common": "ノルウェー"
      },
      "kor": {
        "official": "노르웨이 왕국",
        "common": "노르웨이"
      },
      "nld": {
        "official": "Koninkrijk Noorwegen",
        "common": "Noorwegen"
      },
      "per": {
        "official": "پادشاهی نروژ",
        "common": "نروژ"
      },
      "pol": {
        "official": "Królestwo Norwegii",
        "common": "Norwegia"
      },
      "por": {
        "official": "Reino da Noruega",
        "common": "Noruega"
      },
      "rus": {
        "official": "Королевство Норвегия",
        "common": "Норвегия"
      },
      "slk": {
        "official": "Nórske kráľovstvo",
        "common": "Nórsko"
      },
      "spa": {
        "official": "Reino de Noruega",
        "common": "Noruega"
      },
      "srp": {
        "official": "Краљевина Норвешка",
        "common": "Норвешка"
      },
      "swe": {
        "official": "Konungariket Norge",
        "common": "Norge"
      },
      "tur": {
        "official": "Norveç Krallığı",
        "common": "Norveç"
      },
      "urd": {
        "official": "مملکتِ ناروے",
        "common": "ناروے"
      },
      "zho": {
        "official": "挪威王国",
        "common": "挪威"
      }
    },
    "latlng": [
      62,
      10
    ],
    "landlocked": false,
    "borders": [
      "FIN",
      "SWE",
      "RUS"
    ],
    "area": 323802,
    "demonyms": {
      "eng": {
        "f": "Norwegian",
        "m": "Norwegian"
      },
      "fra": {
        "f": "Norvégienne",
        "m": "Norvégien"
      }
    },
    "flag": "🇳🇴",
    "maps": {
      "googleMaps": "https://goo.gl/maps/htWRrphA7vNgQNdSA",
      "openStreetMaps": "https://www.openstreetmap.org/relation/2978650"
    },
    "population": 5379475,
    "gini": {
      "2018": 27.6
    },
    "fifa": "NOR",
    "car": {
      "signs": [
        "N"
      ],
      "side": "right"
    },
    "timezones": [
      "UTC+01:00"
    ],
    "continents": [
      "Europe"
    ],
    "flags": {
      "png": "https://flagcdn.com/w320/no.png",
      "svg": "https://flagcdn.com/no.svg",
      "alt": "The flag of Norway has a red field with a large white-edged navy blue cross that extends to the edges of the field. The vertical part of this cross is offset towards the hoist side."
    },
    "coatOfArms": {
      "png": "https://mainfacts.com/media/images/coats_of_arms/no.png",
      "svg": "https://mainfacts.com/media/images/coats_of_arms/no.svg"
    },
    "startOfWeek": "monday",
    "capitalInfo": {
      "latlng": [
        59.92,
        10.75
      ]
    },
    "postalCode": {
      "format": "####",
      "regex": "^(\\d{4})$"
    }
  }
]
`

const germanyData = `
[
  {
    "name": {
      "common": "Germany",
      "official": "Federal Republic of Germany",
      "nativeName": {
        "deu": {
          "official": "Bundesrepublik Deutschland",
          "common": "Deutschland"
        }
      }
    },
    "tld": [
      ".de"
    ],
    "cca2": "DE",
    "ccn3": "276",
    "cca3": "DEU",
    "cioc": "GER",
    "independent": true,
    "status": "officially-assigned",
    "unMember": true,
    "currencies": {
      "EUR": {
        "name": "Euro",
        "symbol": "€"
      }
    },
    "idd": {
      "root": "+4",
      "suffixes": [
        "9"
      ]
    },
    "capital": [
      "Berlin"
    ],
    "altSpellings": [
      "DE",
      "Federal Republic of Germany",
      "Bundesrepublik Deutschland"
    ],
    "region": "Europe",
    "subregion": "Western Europe",
    "languages": {
      "deu": "German"
    },
    "translations": {
      "ara": {
        "official": "جمهورية ألمانيا الاتحادية",
        "common": "ألمانيا"
      },
      "bre": {
        "official": "Republik Kevreadel Alamagn",
        "common": "Alamagn"
      },
      "ces": {
        "official": "Spolková republika Německo",
        "common": "Německo"
      },
      "cym": {
        "official": "Federal Republic of Germany",
        "common": "Germany"
      },
      "deu": {
        "official": "Bundesrepublik Deutschland",
        "common": "Deutschland"
      },
      "est": {
        "official": "Saksamaa Liitvabariik",
        "common": "Saksamaa"
      },
      "fin": {
        "official": "Saksan liittotasavalta",
        "common": "Saksa"
      },
      "fra": {
        "official": "République fédérale d'Allemagne",
        "common": "Allemagne"
      },
      "hrv": {
        "official": "Njemačka Federativna Republika",
        "common": "Njemačka"
      },
      "hun": {
        "official": "Német Szövetségi Köztársaság",
        "common": "Németország"
      },
      "ita": {
        "official": "Repubblica federale di Germania",
        "common": "Germania"
      },
      "jpn": {
        "official": "ドイツ連邦共和国",
        "common": "ドイツ"
      },
      "kor": {
        "official": "독일 연방 공화국",
        "common": "독일"
      },
      "nld": {
        "official": "Bondsrepubliek Duitsland",
        "common": "Duitsland"
      },
      "per": {
        "official": "جمهوری فدرال آلمان",
        "common": "آلمان"
      },
      "pol": {
        "official": "Republika Federalna Niemiec",
        "common": "Niemcy"
      },
      "por": {
        "official": "República Federal da Alemanha",
        "common": "Alemanha"
      },
      "rus": {
        "official": "Федеративная Республика Германия",
        "common": "Германия"
      },
      "slk": {
        "official": "Nemecká spolková republika",
        "common": "Nemecko"
      },
      "spa": {
        "official": "República Federal de Alemania",
        "common": "Alemania"
      },
      "srp": {
        "official": "Савезна Република Немачка",
        "common": "Немачка"
      },
      "swe": {
        "official": "Förbundsrepubliken Tyskland",
        "common": "Tyskland"
      },
      "tur": {
        "official": "Almanya Federal Cumhuriyeti",
        "common": "Almanya"
      },
      "urd": {
        "official": "وفاقی جمہوریہ جرمنی",
        "common": "جرمنی"
      },
      "zho": {
        "official": "德意志联邦共和国",
        "common": "德国"
      }
    },
    "latlng": [
      51,
      9
    ],
    "landlocked": false,
    "borders": [
      "AUT",
      "BEL",
      "CZE",
      "DNK",
      "FRA",
      "LUX",
      "NLD",
      "POL",
      "CHE"
    ],
    "area": 357114,
    "demonyms": {
      "eng": {
        "f": "German",
        "m": "German"
      },
      "fra": {
        "f": "Allemande",
        "m": "Allemand"
      }
    },
    "flag": "🇩🇪",
    "maps": {
      "googleMaps": "https://goo.gl/maps/mD9FBMq1nvXUBrkv6",
      "openStreetMaps": "https://www.openstreetmap.org/relation/51477"
    },
    "population": 83240525,
    "gini": {
      "2016": 31.9
    },
    "fifa": "GER",
    "car": {
      "signs": [
        "DY"
      ],
      "side": "right"
    },
    "timezones": [
      "UTC+01:00"
    ],
    "continents": [
      "Europe"
    ],
    "flags": {
      "png": "https://flagcdn.com/w320/de.png",
      "svg": "https://flagcdn.com/de.svg",
      "alt": "The flag of Germany is composed of three equal horizontal bands of black, red and gold."
    },
    "coatOfArms": {
      "png": "https://mainfacts.com/media/images/coats_of_arms/de.png",
      "svg": "https://mainfacts.com/media/images/coats_of_arms/de.svg"
    },
    "startOfWeek": "monday",
    "capitalInfo": {
      "latlng": [
        52.52,
        13.4
      ]
    },
    "postalCode": {
      "format": "#####",
      "regex": "^(\\d{5})$"
    }
  }
]
`

const usaData = `
[
  {
    "name": {
      "common": "United States",
      "official": "United States of America",
      "nativeName": {
        "eng": {
          "official": "United States of America",
          "common": "United States"
        }
      }
    },
    "tld": [
      ".us"
    ],
    "cca2": "US",
    "ccn3": "840",
    "cca3": "USA",
    "cioc": "USA",
    "independent": true,
    "status": "officially-assigned",
    "unMember": true,
    "currencies": {
      "USD": {
        "name": "United States dollar",
        "symbol": "$"
      }
    },
    "idd": {
      "root": "+1",
      "suffixes": [
        "201",
        "202",
        "203",
        "205",
        "206",
        "207",
        "208",
        "209",
        "210",
        "212",
        "213",
        "214",
        "215",
        "216",
        "217",
        "218",
        "219",
        "220",
        "224",
        "225",
        "227",
        "228",
        "229",
        "231",
        "234",
        "239",
        "240",
        "248",
        "251",
        "252",
        "253",
        "254",
        "256",
        "260",
        "262",
        "267",
        "269",
        "270",
        "272",
        "274",
        "276",
        "281",
        "283",
        "301",
        "302",
        "303",
        "304",
        "305",
        "307",
        "308",
        "309",
        "310",
        "312",
        "313",
        "314",
        "315",
        "316",
        "317",
        "318",
        "319",
        "320",
        "321",
        "323",
        "325",
        "327",
        "330",
        "331",
        "334",
        "336",
        "337",
        "339",
        "346",
        "347",
        "351",
        "352",
        "360",
        "361",
        "364",
        "380",
        "385",
        "386",
        "401",
        "402",
        "404",
        "405",
        "406",
        "407",
        "408",
        "409",
        "410",
        "412",
        "413",
        "414",
        "415",
        "417",
        "419",
        "423",
        "424",
        "425",
        "430",
        "432",
        "434",
        "435",
        "440",
        "442",
        "443",
        "447",
        "458",
        "463",
        "464",
        "469",
        "470",
        "475",
        "478",
        "479",
        "480",
        "484",
        "501",
        "502",
        "503",
        "504",
        "505",
        "507",
        "508",
        "509",
        "510",
        "512",
        "513",
        "515",
        "516",
        "517",
        "518",
        "520",
        "530",
        "531",
        "534",
        "539",
        "540",
        "541",
        "551",
        "559",
        "561",
        "562",
        "563",
        "564",
        "567",
        "570",
        "571",
        "573",
        "574",
        "575",
        "580",
        "585",
        "586",
        "601",
        "602",
        "603",
        "605",
        "606",
        "607",
        "608",
        "609",
        "610",
        "612",
        "614",
        "615",
        "616",
        "617",
        "618",
        "619",
        "620",
        "623",
        "626",
        "628",
        "629",
        "630",
        "631",
        "636",
        "641",
        "646",
        "650",
        "651",
        "657",
        "660",
        "661",
        "662",
        "667",
        "669",
        "678",
        "681",
        "682",
        "701",
        "702",
        "703",
        "704",
        "706",
        "707",
        "708",
        "712",
        "713",
        "714",
        "715",
        "716",
        "717",
        "718",
        "719",
        "720",
        "724",
        "725",
        "727",
        "730",
        "731",
        "732",
        "734",
        "737",
        "740",
        "743",
        "747",
        "754",
        "757",
        "760",
        "762",
        "763",
        "765",
        "769",
        "770",
        "772",
        "773",
        "774",
        "775",
        "779",
        "781",
        "785",
        "786",
        "801",
        "802",
        "803",
        "804",
        "805",
        "806",
        "808",
        "810",
        "812",
        "813",
        "814",
        "815",
        "816",
        "817",
        "818",
        "828",
        "830",
        "831",
        "832",
        "843",
        "845",
        "847",
        "848",
        "850",
        "854",
        "856",
        "857",
        "858",
        "859",
        "860",
        "862",
        "863",
        "864",
        "865",
        "870",
        "872",
        "878",
        "901",
        "903",
        "904",
        "906",
        "907",
        "908",
        "909",
        "910",
        "912",
        "913",
        "914",
        "915",
        "916",
        "917",
        "918",
        "919",
        "920",
        "925",
        "928",
        "929",
        "930",
        "931",
        "934",
        "936",
        "937",
        "938",
        "940",
        "941",
        "947",
        "949",
        "951",
        "952",
        "954",
        "956",
        "959",
        "970",
        "971",
        "972",
        "973",
        "975",
        "978",
        "979",
        "980",
        "984",
        "985",
        "989"
      ]
    },
    "capital": [
      "Washington, D.C."
    ],
    "altSpellings": [
      "US",
      "USA",
      "United States of America"
    ],
    "region": "Americas",
    "subregion": "North America",
    "languages": {
      "eng": "English"
    },
    "translations": {
      "ara": {
        "official": "الولايات المتحدة الامريكية",
        "common": "الولايات المتحدة"
      },
      "bre": {
        "official": "Stadoù-Unanet Amerika",
        "common": "Stadoù-Unanet"
      },
      "ces": {
        "official": "Spojené státy americké",
        "common": "Spojené státy"
      },
      "cym": {
        "official": "United States of America",
        "common": "United States"
      },
      "deu": {
        "official": "Vereinigte Staaten von Amerika",
        "common": "Vereinigte Staaten"
      },
      "est": {
        "official": "Ameerika Ühendriigid",
        "common": "Ameerika Ühendriigid"
      },
      "fin": {
        "official": "Amerikan yhdysvallat",
        "common": "Yhdysvallat"
      },
      "fra": {
        "official": "Les états-unis d'Amérique",
        "common": "États-Unis"
      },
      "hrv": {
        "official": "Sjedinjene Države Amerike",
        "common": "Sjedinjene Američke Države"
      },
      "hun": {
        "official": "Amerikai Egyesült Államok",
        "common": "Amerikai Egyesült Államok"
      },
      "ita": {
        "official": "Stati Uniti d'America",
        "common": "Stati Uniti d'America"
      },
      "jpn": {
        "official": "アメリカ合衆国",
        "common": "アメリカ合衆国"
      },
      "kor": {
        "official": "아메리카 합중국",
        "common": "미국"
      },
      "nld": {
        "official": "Verenigde Staten van Amerika",
        "common": "Verenigde Staten"
      },
      "per": {
        "official": "ایالات متحده آمریکا",
        "common": "ایالات متحده آمریکا"
      },
      "pol": {
        "official": "Stany Zjednoczone Ameryki",
        "common": "Stany Zjednoczone"
      },
      "por": {
        "official": "Estados Unidos da América",
        "common": "Estados Unidos"
      },
      "rus": {
        "official": "Соединенные Штаты Америки",
        "common": "Соединённые Штаты Америки"
      },
      "slk": {
        "official": "Spojené štáty Americké",
        "common": "Spojené štáty americké"
      },
      "spa": {
        "official": "Estados Unidos de América",
        "common": "Estados Unidos"
      },
      "srp": {
        "official": "Сједињене Америчке Државе",
        "common": "Сједињене Америчке Државе"
      },
      "swe": {
        "official": "Amerikas förenta stater",
        "common": "USA"
      },
      "tur": {
        "official": "Amerika Birleşik Devletleri",
        "common": "Amerika Birleşik Devletleri"
      },
      "urd": {
        "official": "ریاستہائے متحدہ امریکا",
        "common": "ریاستہائے متحدہ"
      },
      "zho": {
        "official": "美利坚合众国",
        "common": "美国"
      }
    },
    "latlng": [
      38,
      -97
    ],
    "landlocked": false,
    "borders": [
      "CAN",
      "MEX"
    ],
    "area": 9372610,
    "demonyms": {
      "eng": {
        "f": "American",
        "m": "American"
      },
      "fra": {
        "f": "Américaine",
        "m": "Américain"
      }
    },
    "flag": "🇺🇸",
    "maps": {
      "googleMaps": "https://goo.gl/maps/e8M246zY4BSjkjAv6",
      "openStreetMaps": "https://www.openstreetmap.org/relation/148838#map=2/20.6/-85.8"
    },
    "population": 329484123,
    "gini": {
      "2018": 41.4
    },
    "fifa": "USA",
    "car": {
      "signs": [
        "USA"
      ],
      "side": "right"
    },
    "timezones": [
      "UTC-12:00",
      "UTC-11:00",
      "UTC-10:00",
      "UTC-09:00",
      "UTC-08:00",
      "UTC-07:00",
      "UTC-06:00",
      "UTC-05:00",
      "UTC-04:00",
      "UTC+10:00",
      "UTC+12:00"
    ],
    "continents": [
      "North America"
    ],
    "flags": {
      "png": "https://flagcdn.com/w320/us.png",
      "svg": "https://flagcdn.com/us.svg",
      "alt": "The flag of the United States of America is composed of thirteen equal horizontal bands of red alternating with white. A blue rectangle, bearing fifty small five-pointed white stars arranged in nine rows where rows of six stars alternate with rows of five stars, is superimposed in the canton."
    },
    "coatOfArms": {
      "png": "https://mainfacts.com/media/images/coats_of_arms/us.png",
      "svg": "https://mainfacts.com/media/images/coats_of_arms/us.svg"
    },
    "startOfWeek": "sunday",
    "capitalInfo": {
      "latlng": [
        38.89,
        -77.05
      ]
    },
    "postalCode": {
      "format": "#####-####",
      "regex": "^\\d{5}(-\\d{4})?$"
    }
  }
]
`
