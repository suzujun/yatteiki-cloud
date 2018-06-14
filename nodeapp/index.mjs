import apiClient from './api_client'
import fs from 'fs';
import express from 'express';
import body_parser from 'body-parser';

var app = express();
// app.use(express.static('public'));
app.use(body_parser.urlencoded({ extended: true }));
app.use(body_parser.json({limit: '10mb'}));

const port = 3000;

app.get('/', (req, res) => {
  // var data = {name:"suzuki",age:25}
  // res.send(data)

  fs.readFile('./views/index.html', 'utf-8', (err, data) => {
    res.writeHead(200, {'Content-Type': 'text/html'});
    res.write(data);
    res.end();
  });
});

app.get('/ping', (req, res, next) => {
  apiClient.ping((err, body) => {
    if (err) {
      return next(err)
    }
    res.send(body)
  })
});

app.get('/pingdb', (req, res, next) => {
  apiClient.pingDb((err, body) => {
    if (err) {
      return next(err)
    }
    res.send(body)
  })
});

app.get('/todos', (req, res, next) => {
  var limit = parseInt(req.query.limit) || 100
  apiClient.getList(limit, req.query.marker, (err, body) => {
    if (err) {
      return next(err)
    }
    res.send(body)
  })
});

app.get('/todos/:id', (req, res, next) => {
  apiClient.get(req.params.id, (err, body) => {
    if (err) {
      return next(err)
    }
    if (!body) {
      return res.status(404).send({error:`not found todos ${req.params.id}`})
    }
    res.send(body)
  })
});

app.post('/todos', (req, res, next) => {
  if (!req.body || !req.body.title) {
    return res.status(400).send({error:'required title'})
  }
  apiClient.create(req.body.title, (err, body) => {
    if (err) {
      return next(err)
    }
    res.send(body)
  })
});

app.patch('/todos/:id', (req, res, next) => {
  if (!req.body.title) {
    return res.status(400).send({error:'required title'})
  }
  var body = {
    title: req.body.title,
  }
  if ("completed" in req.body) {
    body.completed = req.body.completed
  }
  apiClient.update(req.params.id, body, (err, body) => {
    if (err) {
      return next(err)
    }
    res.send(body)
  })
});

app.delete('/todos/:id', (req, res, next) => {
  apiClient.delete(req.params.id, (err, body) => {
    if (err) {
      return next(err)
    }
    res.send(body)
  })
});

// 404 Not Found 共通
app.use((req, res, next) => {
  console.warn('API Not Found, [%s] %s', req.method, req.url);
  res.status(404).json({error:"not found path"});
});

app.use((err, req, res, next) => {
  if (!err) {
    return next();
  }
  res.status(500).json({error:err});
});

app.listen(port, error => {
  if (error) {
    console.error(error);
  } else {
    console.info('listen: ', port);
  }
});
