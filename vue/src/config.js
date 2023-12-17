// 后端URL
const serverURL = "http://localhost";
const serverPort = 9000;
const captcha = "/captcha";

function getServer() {
  return serverURL + ":" + serverPort;
}

function getCaptcha() {
  return serverURL + ":" + serverPort + captcha;
}

export { getServer, getCaptcha };
