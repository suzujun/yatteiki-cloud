/*global riot, todoStorage */

(function () {
	'use strict';

	// riot.mount('todo', { data: todoStorage.fetch() });
	todoStorage.fetch(function(err, res){
		if (err) {
			alert("fetch failed, err=" + err)
			return
		}
		console.log(">>>>>>>>>>>>>>res", res)
		riot.mount('todo', { data: res });
	})
}());