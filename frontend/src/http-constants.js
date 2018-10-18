import axios from "axios";

// let baseURL;

// if (!process.env.NODE_ENV || process.env.NODE_ENV === "development") {
//   baseURL = "http://localhost:8090/";
// } else {
//   baseURL = "http://api.example.com:8000/";
// }

export const goAPI = axios.create({
  baseURL: "http://localhost:8090/",
  timeout: Number.MAX_SAFE_INTEGER,
  headers: {
    "content-type": "application/x-www-form-urlencoded"
  }
});
