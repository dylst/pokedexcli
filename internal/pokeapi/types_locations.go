package pokeapi

/*
map[count:1089 next:https://pokeapi.co/api/v2/location-area?offset=20&limit=20 previous:<nil>
results:[map[name:canalave-city-area url:https://pokeapi.co/api/v2/location-area/1/]
		 map[name:eterna-city-area url:https://pokeapi.co/api/v2/location-area/2/]
		 map[name:pastoria-city-area url:https://pokeapi.co/api/v2/location-area/3/]
		 map[name:sunyshore-city-area url:https://pokeapi.co/api/v2/location-area/4/]
		 map[name:sinnoh-pokemon-league-area url:https://pokeapi.co/api/v2/location-area/5/]
		 map[name:oreburgh-mine-1f url:https://pokeapi.co/api/v2/location-area/6/]
		 map[name:oreburgh-mine-b1f url:https://pokeapi.co/api/v2/location-area/7/]
		 map[name:valley-windworks-area url:https://pokeapi.co/api/v2/location-area/8/]
		 map[name:eterna-forest-area url:https://pokeapi.co/api/v2/location-area/9/]
		 map[name:fuego-ironworks-area url:https://pokeapi.co/api/v2/location-area/10/]
		 map[name:mt-coronet-1f-route-207 url:https://pokeapi.co/api/v2/location-area/11/]
		 map[name:mt-coronet-2f url:https://pokeapi.co/api/v2/location-area/12/]
		 map[name:mt-coronet-3f url:https://pokeapi.co/api/v2/location-area/13/]
		 map[name:mt-coronet-exterior-snowfall url:https://pokeapi.co/api/v2/location-area/14/]
		 map[name:mt-coronet-exterior-blizzard url:https://pokeapi.co/api/v2/location-area/15/]
		 map[name:mt-coronet-4f url:https://pokeapi.co/api/v2/location-area/16/]
		 map[name:mt-coronet-4f-small-room url:https://pokeapi.co/api/v2/location-area/17/]
		 map[name:mt-coronet-5f url:https://pokeapi.co/api/v2/location-area/18/]
		 map[name:mt-coronet-6f url:https://pokeapi.co/api/v2/location-area/19/]
		 map[name:mt-coronet-1f-from-exterior url:https://pokeapi.co/api/v2/location-area/20/]
		]]
*/
type RespLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}