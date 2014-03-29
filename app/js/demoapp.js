$(document).foundation();
$(document).ready(function() {
	$booklink = $(".book");
	$booklink.each(function() {
		$this = $(this);
		var show = $this.data("value");
		window.console.log(show);
		$this.click(function() {
			booker(show);
		})
	});
});
var booker = function (id) {
	$.getJSON("/api/books/"+id, function (book) {
		window.console.log(book.data.chapter);
		$demoapp = $('.demoapp');
		$demoapp.empty();
		$demoapp.append("<h1>"+book.meta.title+"</h1>");
		$demoapp.append("<h5>"+book.meta.creator+"</h5>");
		for (var p in book.data.part) {
			for (var chap in book.data.part[p].chapter) {
				$demoapp.append("<p>"+book.data.part[p].chapter[chap].title+"</p>");
				for (var par in book.data.part[p].chapter[chap].text) {
					$demoapp.append("<p>"+book.data.part[p].chapter[chap].text[par]+"</p>");
				};
			};
		}
		for (var chap in book.data.chapter) {
			$demoapp.append("<p>"+book.data.chapter[chap].title+"</p>");
			for (var par in book.data.chapter[chap].text) {
				$demoapp.append("<p>"+book.data.chapter[chap].text[par]+"</p>");
			};
		};
	});

};