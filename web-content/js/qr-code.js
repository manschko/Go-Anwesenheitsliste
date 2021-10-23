var targetHost = "http://localhost";
var expirationTime = 0;

function startRequest (method, link, responseMethod, data) {
  var requester = new XMLHttpRequest();
  requester.onreadystatechange = responseMethod

  requester.open(method, location, true);

  if (data !== null) {
    requester.send(data);
  } else {
    requester.send();
  }
}

function getLocation() {
  return document.getElementById(
    "head"
  ).innerText.replace("Anwesenheitsliste ", "");
}

function setTokenExpirationTime() {
  var link = targetHost + "/get-token-expiration-time";
  var method = function(event) {
    if (event.readyState === 4) {
      if (event.status === 200) {
        //TODO:

        setTimeout(updateLocationToken, 20000);
      }
    }
  };

  startRequest("GET", link, method, null);
}

function updateLocationToken() {
  var location = getLocation();
  var link = targetHost + "/get-location-token?location=" + location;

  var response = function(event) {
    if (event.readyState === 4) {
      if (event.status === 200) {
        //TODO:

        setTimeout(updateLocationToken, expirationTime);
      }
    }
  };

  startRequest("GET", link, response, null);
}

setTokenExpirationTime();
