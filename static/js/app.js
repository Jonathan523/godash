// webSocket
const WsType = { Weather: 0, System: 1 };
const apiBase = window.location.origin + "/api";
let timeOut = 1;
connect();

// weather elements
const weatherIcon = document.getElementById("weatherIcon");
const weatherTemp = document.getElementById("weatherTemp");
const weatherDescription = document.getElementById("weatherDescription");
const weatherHumidity = document.getElementById("weatherHumidity");
const weatherSunrise = document.getElementById("weatherSunrise");
const weatherSunset = document.getElementById("weatherSunset");

// system elements
const systemCpu = document.getElementById("systemCpu");
const systemRamPercentage = document.getElementById("systemRamPercentage");
const systemDiskPercentage = document.getElementById("systemDiskPercentage");
const systemUptime = document.getElementById("systemUptime");

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
  weatherIcon.setAttribute("xlink:href", "#" + parsed.weather[0].icon);
  weatherTemp.innerText = parsed.main.temp;
  weatherDescription.innerText = parsed.weather[0].description;
  weatherHumidity.innerText = parsed.main.humidity;
  weatherSunrise.innerText = parsed.sys.str_sunrise;
  weatherSunset.innerText = parsed.sys.str_sunset;
}

function replaceSystem(parsed) {
  systemCpu.innerText = parsed.cpu.percentage;
  systemRamPercentage.innerText = parsed.ram.percentage;
  systemDiskPercentage.innerText = parsed.disk.percentage;
  systemUptime.innerText = parsed.server_uptime;
}
