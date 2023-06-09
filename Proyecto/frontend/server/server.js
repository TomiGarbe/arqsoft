import express from 'express';

const app = express();

const fetch = (...args) => 
    import('node-fetch').then(({default: fetch}) => fetch(...args));

const baseURL = "http://localhost:8090";

app.get("/api/cliente/email/:email", async function(req, res) {
    res.set("Access-Control-Allow-Origin", "*");
    const url = `${baseURL}/cliente/email/${req.params.email}`;
    const options = {method: 'GET'};
    try {
        let response = await fetch(url, options);
        let status = response.status;
        response = await response.json();
        res.status(status).json(response);
    } catch (err) {
        console.log(err);
        res.status(500).json({msg: 'Internal Server Error'});
    }
});

/*app.post("/api/cliente", async function(req, res) {
    res.set("Access-Control-Allow-Origin", "*");
    const formData = req;
    const url = `${baseURL}/cliente`;
    const options = { method: 'POST', body: JSON.stringify(formData)};
    try {
      let response = await fetch(url, options);
      let status = response.status;
      response = await response.json();
      res.status(status).json(response);
    } catch (err) {
      console.log(err);
      res.status(500).json({ msg: 'Internal Server Error' });
    }
  });*/
  
/*Casi funciona
app.post("/api/cliente/:formData", async function(req, res) {
    res.set("Access-Control-Allow-Origin", "*");
    const formData = req.params.formData;
    const url = `${baseURL}/cliente`;
    const options = {
        method: 'POST',
        body: JSON.stringify(formData)
        //body: {
        //    "id": 6,
        //    "name": "Tomi",
        //    "last_name": "Garbe",
        //    "username": "tomig",
        //    "password": "12345",
        //    "email": "tomig@123"
        //  }
      };
    try {
        let response = await fetch(url, options);
        let status = response.status;
        response = await response.json();
        res.status(status).json(response);
    } catch (err) {
        console.log(err);
        res.status(500).json({msg: 'Internal Server Error'});
    }
});*/

app.post("/api/cliente", async function(req, res) {
    res.set("Access-Control-Allow-Origin", "*");
    const formData = req.body; // Obtén el cuerpo de la solicitud POST
    const url = `${baseURL}/cliente`;
    const options = {
      method: 'POST',
      body: JSON.stringify({formData})
    };
    try {
      let response = await fetch(url, options);
      let status = response.status;
      response = await response.json();
      res.status(status).json(response);
    } catch (err) {
      console.log(err);
      res.status(500).json({ msg: 'Internal Server Error' });
    }
  });
  

app.listen(5001, () => {
    console.log("Server started on port 5001");
});