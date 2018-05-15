'use strict';
const app = new (require('express'))();
const port = 3000;

app.get('/', (req, res) => {
  var data = {name:"suzuki",age:25}
  res.send(data)
});

app.listen(port, error => {
  if (error) {
    console.error(error);
  } else {
    console.info('listen: ', port);
  }
});

