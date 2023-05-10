import React from "react";

const LogoutButton = () => {
  async function logoutAuthService() {
    window.location.href = "http://localhost:8080/logout"
  }
  return (
    <button onClick={() => logoutAuthService()}>
      Log Out
    </button>
  );
};

export default LogoutButton;