'use strict';

new Vue({
  el: '#vueapp',
  data: {
    keypressed: ""
  },
  methods: {
    ShowKeyCode: function (event) {
      this.keypressed = event.keyCode;
    }
  }
});
