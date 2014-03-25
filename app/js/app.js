$(document).foundation();
$(document).ready(function () {
  $.getJSON("/api/books/",
    function (books) {
      var tr;
      $table = $('table');
      for (var i = 0; i < books.length; i++) {
        $table.append("<tr></tr>");
        $tr = $('tr').last();
        $tr.append("<td class='id'>" + books[i].id + "</td>");
        $tr.append("<td class='author'>" + books[i].author + "</td>");
        $tr.append("<a class='link' href='/api/books/"+books[i].id+"'>" + books[i].title + "</a>");
        $tr.append("</tr>");
        $('.link').last().wrap("<td class='title'></td>");
      }
    });
});