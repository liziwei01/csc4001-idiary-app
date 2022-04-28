import http from "./httpService";

const apiUrl = "http://jsonplaceholder.typicode.com";
const apiEndpoint = apiUrl + "/users";

export function register(user) {
  return http.post(apiEndpoint, {
    email: user.email,
    password: user.password,
    name: user.username,
  });
}
