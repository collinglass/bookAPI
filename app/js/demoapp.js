$(document).foundation();
$(document).ready(function () {
  $.getJSON("/api/books/1", function (book) {
      window.console.log(book);
      $demoapp = $('.demoapp');
      $demoapp.append("<h1>"+book.meta.title[0]+"</h1>");
      $demoapp.append("<h5>"+book.meta.creator[0]+"</h5>");
      $demoapp.append("<p>"+book.data.chapter[1].title+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[1]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[2]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[3]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[4]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[5]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[7]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[8]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[9]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[10]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[11]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[12]+"</p>");
      $demoapp.append("<p>"+book.data.chapter[1].text[13]+"</p>");
    });
});