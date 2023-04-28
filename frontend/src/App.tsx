import React, {useState} from 'react';
import './App.css';
import Header from "./components/Header";
import ShortenURL from "./components/ShortenURL";
import BasicTable from "./components/BasicTable";
import {ShortenURLS} from "./models/models";

function App() {
  const [allURLS, setAllURLS] = useState<ShortenURLS[]>([])

  return (
    <div>
      <Header />
      <ShortenURL setAllURLS={setAllURLS}/>
      {
        allURLS.length !== 0 && <BasicTable allURLS={allURLS}/>
      }
    </div>
  );
}

export default App;
