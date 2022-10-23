import { select } from "https://cdn.skypack.dev/d3-selection@3";
import { timeDay } from "https://cdn.skypack.dev/d3-time@3";

const WsType = { Weather: 0, System: 1 };
const apiBase = window.location.origin + "/api";
let timeOut = 1;
connect();

function connect() {
  let ws = new WebSocket(apiBase.replace("http", "ws") + "/system/ws");
  ws.onopen = () => {
    console.log("WebSocket is open.");
    timeOut = 1;
  };
  ws.onmessage = (event) => handleMessage(JSON.parse(event.data));
  ws.onerror = () => ws.close();
  ws.onclose = () => {
    console.log("WebSocket is closed. Reconnect will be attempted in " + timeOut + " second.");
    setTimeout(() => connect(), timeOut * 1000);
    timeOut += 1;
  };
}

function handleMessage(parsed) {
  if (parsed.ws_type === WsType.Weather) replaceWeather(parsed.message);
  else if (parsed.ws_type === WsType.System) replaceSystem(parsed.message);
}

function replaceWeather(parsed) {
  select("#weatherIcon").attr("xlink:href", "#" + parsed.weather[0].icon);
  select("#weatherTemp").text(parsed.main.temp);
  select("#weatherDescription").text(parsed.weather[0].description);
  select("#weatherHumidity").text(parsed.main.humidity);
  select("#weatherSunrise").text(parsed.sys.str_sunrise);
  select("#weatherSunset").text(parsed.sys.str_sunset);
}

function replaceSystem(parsed) {
  select("#systemCpu").text(parsed.cpu.percentage);
  select("#systemRamPercentage").text(parsed.ram.percentage);
  select("#systemDiskPercentage").text(parsed.disk.percentage);
  select("#systemUptime").text(parsed.server_uptime);
}
