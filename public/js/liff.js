$(document).ready(function () {
  var liffId = "1655211409-WNGnDO0J";
  initializeLiff(liffId);
})

function initializeLiff(liffId) {
  liff
      .init({
          liffId: liffId
      })
      .then(() => {
          if (!liff.isInClient() && !liff.isLoggedIn()) {
              window.alert("LINEアカウントにログインしてください。");
              liff.login();
          }
      })
      .catch((err) => {
          console.log('LIFF Initialization failed ', err);
      });
}

function sendMessage(text) {
  if (liff.isInClient()) {
      sendMessages(text);
  } else {
      shareTargetPicker(text);
  }
}

function sendMessages(text) {
  liff.sendMessages([{
      'type': 'text',
      'text': text
  }]).then(function () {
      liff.closeWindow();
  }).catch(function (error) {
      window.alert('Failed to send message sendMessages ' + error);
  });
}

function shareTargetPicker(text) {
  liff.shareTargetPicker([{
      'type': 'text',
      'text': text
  }]).catch(function (error) {
      window.alert('Failed to send message shareTargetPicker ' + error);
  });
}