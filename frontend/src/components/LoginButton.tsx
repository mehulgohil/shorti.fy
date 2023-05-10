import React from "react";

  const LoginButton = () => {

  async function loginAuthService() {
    window.location.href = "http://localhost:8080/login"
  }

  return <button onClick={() => loginAuthService()}>Log In</button>;
};

export default LoginButton;