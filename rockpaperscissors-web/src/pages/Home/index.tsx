import React from "react";
import { Col, Container, Row } from "react-bootstrap";
import StartGameForm from "src/components/StartGameForm";

const Home = () => {
  // useEffect(() => {
  //   connect((msg, err) => {
  //     if (msg) {
  //       window.alert(JSON.stringify(msg));
  //     }
  //     if (err) {
  //       window.alert("Errored");
  //     }
  //   });

  //   return () => {
  //     window.alert("CLOSED");
  //     close();
  //   };
  // }, []);

  // const sendPing = () => {
  //   sendMsg("ping");
  // };

  // const handleCreateGame = () => {
  //   axios
  //     .post(`${process.env.REACT_APP_API_URL}/games`, { rounds: 1 })
  //     .then((response) => {
  //       console.log("Response: ", response);
  //       window.alert(JSON.stringify(response.data));
  //     })
  //     .catch((error) => {
  //       console.log("Errored: ", error);
  //     });
  // };

  // const handleJoinGame = () => {};

  return (
    <Container>
      {/* <Button onClick={handleCreateGame}>Create Game</Button>
      <Button onClick={handleJoinGame}>Join Game</Button> */}
      <Row>
        <Col>
          <StartGameForm startType="create" />
        </Col>
        <Col>
          <StartGameForm startType="join" />
        </Col>
      </Row>
    </Container>
  );
};

export default Home;
