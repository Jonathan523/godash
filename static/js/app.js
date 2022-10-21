const WsType = { Weather: 0, System: 1 };
let socket = new WebSocket(window.location.origin.replace("http", "ws") + "/api/system/ws");
const weatherIcon = document.getElementById("weatherIcon");
const weatherTemp = document.getElementById("weatherTemp");

socket.onmessage = (event) => {
  const parsed = JSON.parse(event.data);
  if (parsed.ws_type === WsType.Weather) {
    const weather = parsed.message;
    weatherIcon.setAttribute("xlink:href", "#" + weather.weather[0].icon);
    weatherTemp.innerHTML = parsed.message.main.temp + " " + parsed.message.units;
  }
};
