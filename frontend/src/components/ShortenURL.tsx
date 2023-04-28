import {Button, FormControl, Grid, TextField} from "@mui/material";
import React, {useState} from "react";
import {ShortenURLS} from "../models/models";

interface ShortenURLProps {
  setAllURLS: React.Dispatch<React.SetStateAction<ShortenURLS[]>>
}

function ShortenURL(props: ShortenURLProps) {
  const [longURL, setLongURL] = useState<string>();
  const [email, setEmail] = useState<string>();

  const onFieldValueChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>, fieldType: string) => {
    switch (fieldType) {
      case "LongURL":
        setLongURL(e.target.value)
        break
      case "Email":
        setEmail(e.target.value)
        break
    }
  }

  async function submitShortenURL() {
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        "long_url": longURL,
        "user_email": email
      })
    };
    const response = await fetch("http://a585e18fb8f114857badcd6d85868d49-1713447191.ap-south-1.elb.amazonaws.com/v1/shorten", requestOptions);

    if(response.ok) {
      const data = JSON.stringify(await response.json());
      const repo = JSON.parse(data)
      props.setAllURLS(prevState => [...prevState, {LongURL: repo.long_url, ShortURL: repo.short_url}])
    } else {
      console.error(JSON.stringify(await response.json()))
      alert("error shortening url")
    }
  }

  return (
    <FormControl style={{marginLeft: "20px"}}>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <TextField id="standard-basic" label="Long URL" variant="standard" onChange={(e) => onFieldValueChange(e, "LongURL")}/>
        </Grid>
        <Grid item xs={12}>
          <TextField id="standard-basic" label="Email" variant="standard" onChange={(e) => onFieldValueChange(e, "Email")}/>
        </Grid>
      </Grid>
      <Button variant="contained" style={{marginTop: "10px"}} onClick={submitShortenURL}>Submit</Button>
    </FormControl>
  )
}

export default ShortenURL