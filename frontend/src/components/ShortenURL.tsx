import {Button, FormControl, Grid, TextField} from "@mui/material";
import React, {useState} from "react";

function ShortenURL() {
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
    const response = await fetch('https://a585e18fb8f114857badcd6d85868d49-1713447191.ap-south-1.elb.amazonaws.com/v1/shorten', requestOptions);
    const data = await response.json();
    alert(data)
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