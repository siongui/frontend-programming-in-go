var textToUrlMapping = {
  "世尊譯詞的探討": "http://agama.buddhason.org/book/bb/bb21.htm",
};

function TextToLink(elm) {
  elm.innerHTML = elm.innerHTML.replace(/〈(.+)〉/g, function(text, str1) {
    if (textToUrlMapping.hasOwnProperty(str1)) {
      return '〈<a href="'+ textToUrlMapping[str1] +
             '" target="_blank">' + str1 +
             '</a>〉';
    }
    return str1;
  });
}

var t = document.querySelector(".line-block");
TextToLink(t);
