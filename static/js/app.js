import { select } from "https://cdn.skypack.dev/d3-selection@3";

const WsType = { Weather: 0, System: 1 };
const apiBase = window.location.origin + "/api";
let socket = new WebSocket(apiBase.replace("http", "ws") + "/system/ws");

socket.onmessage = (event) => {
  const parsed = JSON.parse(event.data);
  if (parsed.ws_type === WsType.Weather) {
    select("#weatherIcon").attr("xlink:href", "#" + parsed.message.weather[0].icon);
    select("#weatherTemp").text(parsed.message.main.temp);
    select("#weatherDescription").text(parsed.message.weather[0].description);
    select("#weatherHumidity").text(parsed.message.main.humidity);
    select("#weatherSunrise").text(parsed.message.sys.str_sunrise);
    select("#weatherSunset").text(parsed.message.sys.str_sunset);
  }
};
