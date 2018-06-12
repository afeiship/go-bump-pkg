(function () {

  var global = global || this || self || window;
  var nx = global.nx || require('next-js-core2');

  nx.auctionTime = function (inStatus, inIsLeave) {
    switch (true) {
      case inStatus === 'N':
        return { key: 'actualStartTime', value: '开拍时间' };
      case inStatus === 'A':
        return { key: 'remainSeconds', value: '倒计时' };
      case inStatus === 'F':
        return { key: 'actualEndTime', value: '截拍时间' };
    }
  };

  if (typeof module !== 'undefined' && module.exports) {
    module.exports = nx.auctionTime;
  }

}());
