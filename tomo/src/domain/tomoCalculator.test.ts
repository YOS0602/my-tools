import { calculateToMo } from "./tomoCalculator";

it.each([
  [
    "全て「まったく違う」で答えた場合、0点が返ること",
    {
      play: 1,
      purpose: 1,
      potential: 1,
      emotionalPressure: 1,
      economicPressure: 1,
      inertia: 1,
    },
    0,
  ],
  [
    "全て「まったくその通りである」と答えた場合、0点が返ること",
    {
      play: 7,
      purpose: 7,
      potential: 7,
      emotionalPressure: 7,
      economicPressure: 7,
      inertia: 7,
    },
    0,
  ],
  [
    "直接的動機がMAXかつ間接的動機がMINである場合、100点が返ること",
    {
      play: 7,
      purpose: 7,
      potential: 7,
      emotionalPressure: 1,
      economicPressure: 1,
      inertia: 1,
    },
    100,
  ],
  [
    "直接的動機がMINかつ間接的動機がMAXである場合、-100点が返ること",
    {
      play: 1,
      purpose: 1,
      potential: 1,
      emotionalPressure: 7,
      economicPressure: 7,
      inertia: 7,
    },
    -100,
  ],
  [
    "すべて真ん中の4で答えた場合、0点が返ること",
    {
      play: 4,
      purpose: 4,
      potential: 4,
      emotionalPressure: 4,
      economicPressure: 4,
      inertia: 4,
    },
    0,
  ],
])("%s", (_, input, expected) => {
  expect(calculateToMo(input)).toBe(expected);
});

// TODO https://heart-quake.com/assessment/tomo_score.php
// もうちょいテスト増やしても良いかも
