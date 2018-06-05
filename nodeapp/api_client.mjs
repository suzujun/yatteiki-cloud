import request from 'superagent';

class ApiClient {

  constructor() {
    this.baseUrl = "localhost:8080/api/v1"
    this.timeout = 3000
  }

  ping(callback) {
    request
      .get(this.baseUrl+'/ping')
      .timeout(this.timeout)
      .end(this._checkResponse(callback))
  }

  pingDb(callback) {
    request
      .get(this.baseUrl+'/pingdb')
      .timeout(this.timeout)
      .end(this._checkResponse(callback))
  }

  getList(limit, marker, callback) {
    request
      .get(this.baseUrl+'/todos')
      .query({limit,marker})
      .timeout(this.timeout)
      .end(this._checkResponse(callback))
  }

  get(id, callback) {
    request
      .get(this.baseUrl+`/todos/${id}`)
      .timeout(this.timeout)
      .end(this._checkResponse(callback))
  }

  create(note, callback) {
    request
      .post(this.baseUrl+'/todos')
      .send({note})
      .timeout(this.timeout)
      .end(this._checkResponse(callback))
  }

  update(id, note, callback) {
    request
      .put(this.baseUrl+`/todos/${id}`)
      .send({note})
      .timeout(this.timeout)
      .end(this._checkResponse(callback))
  }

  delete(id, callback) {
    request
      .delete(this.baseUrl+`/todos/${id}`)
      .timeout(this.timeout)
      .end(this._checkResponse(callback))
  }

  _checkResponse(callback) {
    return (err, res) => {
      let body = res && res.body;
      let statusCode = res && res.status || 0;
      if (statusCode == 404) {
        return callback(null, null)
      }
      if (err || statusCode < 200 || 300 <= statusCode) {
        return callback(`failed to request err=${err}, statusCode=${statusCode}`);
      }
      callback(null, body);
    }
  }
}

export default new ApiClient();