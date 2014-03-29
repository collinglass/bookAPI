$(document).foundation();
$(document).ready(function() {
	booker(1);
});

var booker = function (id) {
	$.getJSON("/api/books/"+id, function (book) {
		window.console.log(book.data.chapter);
		$demoapp = $('.demoapp');
		$demoapp.append("<h1>"+book.meta.title+"</h1>");
		$demoapp.append("<h5>"+book.meta.creator+"</h5>");
		for (var p in book.data.part) {
			for (var chap in book.data.part[p].chapter) {
				window.console.log("hello");
				$demoapp.append("<p>"+book.data.part[p].chapter[chap].title+"</p>");
				for (var par in book.data.part[p].chapter[chap].text) {
					$demoapp.append("<p>"+book.data.part[p].chapter[chap].text[par]+"</p>");
				};
			};
		}
		for (var chap in book.data.chapter) {
			window.console.log("hello");
			$demoapp.append("<p>"+book.data.chapter[chap].title+"</p>");
			for (var par in book.data.chapter[chap].text) {
				$demoapp.append("<p>"+book.data.chapter[chap].text[par]+"</p>");
			};
		};
	});

};