chrome.tabs.query({'active': true, 'currentWindow': true}, function (tabs) {
    var tab = tabs[0];
    document.write(tab.url);
});
