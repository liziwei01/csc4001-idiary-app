import http from "./httpService";

const apiEndpoint = "http://localhost:8080/user/login";
export function register(user) {
  return http.post(apiEndpoint, {
    email: user.email,
    password: user.password,
    name: user.username,
  });
}

export function login(user) {
  const apiUrl = `http://localhost:8080/user/login`;
  return http.post(apiUrl, {
    password: user.password,
    email: user.email,
  });
}

export function send_email(user) {
  return http.post("http://localhost:8080/email/verificationCode", {
    email: user.email,
  });
}

export function send_problem(user) {
  return http.post(apiEndpoint, {
    username: user.username,
    answer1: user.answer1,
    answer2: user.answer2,
    answer3: user.answer3,
  });
}
