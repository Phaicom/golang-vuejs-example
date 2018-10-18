import axios from "axios";

let baseURL;

if (!process.env.NODE_ENV || process.env.NODE_ENV === "development") {
  baseURL = "http://localhost:8090/";
} else {
  baseURL = "http://api.example.com:8000/";
}

var goAPI = axios.create({
  baseURL: baseURL,
  timeout: Number.MAX_SAFE_INTEGER
  // headers: { "X-Custom-Header": "CustomHeader2" }
});

export default goAPI;
