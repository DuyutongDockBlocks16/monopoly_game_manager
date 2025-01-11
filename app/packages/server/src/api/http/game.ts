import express from "express";
import gameStore from "../../gameStore";
import { ICreateGameRequest, IJoinGameRequest, IJoinGameResponse } from "../dto";
import { getAllProperties } from "./propertyService";


export const subRoute = "/api/game";

const router = express.Router();

// Create a new game
router.post("/", (req, res) => {
  const { name } = req.body as ICreateGameRequest;

  const { gameId, userToken, playerId } = gameStore.createGame(name);

  const response: IJoinGameResponse = { gameId, userToken, playerId };
  res.json(response);
  res.end();
});

router.get('/', (req, res) => {
  res.send('Server is running');
});

// Join a game
router.post("/:gameId", (req, res) => {
  const { gameId } = req.params;
  const { name } = req.body as IJoinGameRequest;

  if (!gameStore.doesGameExist(gameId)) {
    res.status(404).send("Game does not exist");
  } else if (!gameStore.getGame(gameId).isGameOpen()) {
    res.status(403).send("Game is not open");
  } else {
    const game = gameStore.getGame(gameId);
    const { userToken, playerId } = game.addPlayer(name);

    const response: IJoinGameResponse = { gameId, userToken, playerId };
    res.json(response);
  }

  res.end();
});

// Get game status
router.get("/:gameId", (req, res) => {
  const { gameId } = req.params;
  const userToken = req.get("Authorization");

  if (userToken === undefined) {
    res.status(401).send("Authorization not supplied");
  } else if (!gameStore.doesGameExist(gameId)) {
    res.status(404).send("Game does not exist");
  } else if (!gameStore.getGame(gameId).isUserInGame(userToken)) {
    res.status(401).send("You are not permitted to make this operation");
  } else {
    const game = gameStore.getGame(gameId);
    const state = game.getGameState();
    res.json(state);
  }

  res.end();
});

// curl -X GET http://localhost:3000/properties
router.get("/properties", async (req, res) => {
  try {
    // 调用获取所有属性的函数
    const properties = await getAllProperties();
    res.status(200).json(properties); // 返回获取到的属性
  } catch (error) {
    console.error('Error fetching properties:', error);
    res.status(500).json({ error: 'Internal server error' }); // 错误处理
  }
});

export default router;
