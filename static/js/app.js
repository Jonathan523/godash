const WsType = { Weather: 0, System: 1 };
const apiBase = window.location.origin + "/api";
let socket = new WebSocket(apiBase.replace("http", "ws") + "/ws");
const weatherIcon = document.getElementById("weatherIcon");
const weatherTemp = document.getElementById("weatherTemp");
const weatherDescription = document.getElementById("weatherDescription");
const weatherHumidity = document.getElementById("weatherHumidity");
const weatherSunrise = document.getElementById("weatherSunrise");
const weatherSunset = document.getElementById("weatherSunset");

socket.onmessage = (event) => {
  const parsed = JSON.parse(event.data);
  if (parsed.ws_type === WsType.Weather) {
    weatherIcon.setAttribute("xlink:href", "#" + parsed.message.weather[0].icon);
    weatherTemp.innerHTML = parsed.message.main.temp + " " + parsed.message.units;
    weatherDescription.innerHTML = parsed.message.weather[0].description;
    weatherHumidity.innerHTML = parsed.message.main.humidity + "%";
    weatherSunrise.innerHTML = parsed.message.sys.str_sunrise;
    weatherSunset.innerHTML = parsed.message.sys.str_sunset;
  }
};
