import React, {useState} from 'react';
import './App.css';
import Header from "./components/Header";
import ShortenURL from "./components/ShortenURL";
import BasicTable from "./components/BasicTable";
import {ShortenURLS} from "./models/models";
import LoginButton from "./components/LoginButton";
import LogoutButton from "./components/Logout";

function App() {
  const [allURLS, setAllURLS] = useState<ShortenURLS[]>([])

  return (
    <div>
      <Header />
      <ShortenURL setAllURLS={setAllURLS}/>
      {
        allURLS.length !== 0 && <BasicTable allURLS={allURLS}/>
      }
      <LoginButton />
      <LogoutButton />
    </div>
  );
}

export default App;
