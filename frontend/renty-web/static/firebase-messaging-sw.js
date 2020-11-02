// [START initialize_firebase_in_sw]
// Give the service worker access to Firebase Messaging.
// Note that you can only use Firebase Messaging here, other Firebase libraries
// are not available in the service worker.
importScripts("https://www.gstatic.com/firebasejs/7.11.0/firebase-app.js");
importScripts(
  "https://www.gstatic.com/firebasejs/7.11.0/firebase-messaging.js"
);

// Initialize the Firebase app in the service worker by passing in the
// messagingSenderId.
firebase.initializeApp({
  apiKey: "AIzaSyBIKMKs-U7gH5HhQSnM8QYveC4YovcGjvI",
  authDomain: "renty-firebase.firebaseapp.com",
  databaseURL: "https://renty-firebase.firebaseio.com",
  projectId: "renty-firebase",
  storageBucket: "renty-firebase.appspot.com",
  messagingSenderId: "500932383497",
  appId: "1:500932383497:web:52f47353b55a7a3e80b9d8",
  measurementId: "G-SVEBS09F96"
});

// Retrieve an instance of Firebase Messaging so that it can handle background
// messages.
const messaging = firebase.messaging();
// [END initialize_firebase_in_sw]

// If you would like to customize notifications that are received in the
// background (Web app is closed or not in browser focus) then you should
// implement this optional method.
// [START background_handler]
messaging.setBackgroundMessageHandler(function(payload) {
  console.log(
    "[firebase-messaging-sw.js] Received background message ",
    payload
  );
  // Customize notification here
  const notificationTitle = "Background Message Title";
  const notificationOptions = {
    body: "Background Message body."
  };

  return self.registration.showNotification(
    notificationTitle,
    notificationOptions
  );
});
// [END background_handler]
