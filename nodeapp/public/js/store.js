(function (exports) {

	'use strict';

	var STORAGE_KEY = 'todos-riotjs';

	exports.todoStorage = {
		fetch: function (callback) {
			// return JSON.parse(localStorage.getItem(STORAGE_KEY) || '[]');
			axios.get('/todos')
			.then(function (res) {
				console.log(res);
				var datas = res && res.data && res.data.data || []
				// datas = datas.map(function(data){
				// 	return {
				// 		id: data.id,
				// 		title: data.note,
				// 		completed: false,
				// 		editing: false
				// 	}
				// })
				callback(null, datas)
			})
			.catch(function (error) {
				console.log(error);
				callback(error, null)
			});
		},
		create: function (value, callback) {
			// localStorage.setItem(STORAGE_KEY, JSON.stringify(todos));
			axios.post('/todos', {title: value})
			.then(function (res) {
				var {insertedId} = res && res.data
				if (callback) {
					callback(null, insertedId)
				}
			})
			.catch(function (error) {
				console.log(error);
				if (callback) {
					callback(error, null)
				}
			});
		},
		update: function (todo, callback) {
			// localStorage.setItem(STORAGE_KEY, JSON.stringify(todos));
			axios.patch('/todos/'+todo.id, todo)
			.then(function (res) {
				console.log(res);
				if (callback) {
					callback(null, res)
				}
			})
			.catch(function (error) {
				console.log(error);
				if (callback) {
					callback(error, null)
				}
			});
		},
		delete: function (id, callback) {
			// localStorage.setItem(STORAGE_KEY, JSON.stringify(todos));
			axios.delete('/todos/'+id)
			.then(function (res) {
				console.log(res);
				if (callback) {
					callback(null, res)
				}
			})
			.catch(function (error) {
				console.log(error);
				if (callback) {
					callback(error, null)
				}
			});
		}
	};

})(window);