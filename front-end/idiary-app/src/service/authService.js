export function getCurrentUser() {
  try {
    const email = localStorage.getItem("token");
    return email;
  } catch (ex) {
    return null;
  }
}

export default {
  getCurrentUser,
};
