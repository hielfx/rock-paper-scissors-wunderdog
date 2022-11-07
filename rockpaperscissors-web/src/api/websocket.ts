//TODO: CHANGE THIS URL TO USE .env
console.log("process.env.REACT_APP_WS_URL: ", process.env.REACT_APP_WS_URL);

type WebsocketCbFunction = (
  msg: MessageEvent<string> | null,
  err?: Event
) => void;

const socket = new WebSocket("ws://localhost:4040/ws");

const close = () => {
  socket.close();
};

const connect = (cb?: WebsocketCbFunction) => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = (msg) => {
    console.log("Message received: ", msg);
    if (cb) {
      cb(msg);
    }
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
    if (cb) {
      cb(null, error);
    }
  };
};

const sendMsg = (msg: string) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg, close };
