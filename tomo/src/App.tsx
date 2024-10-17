import {
  type ChangeEvent,
  type FormEventHandler,
  useCallback,
  useState,
} from "react";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { tomoQuestions } from "./libs/tomoQuestions";
import { type MotivationKey } from "./libs/tomo.type";
import {
  QuestionSection,
  type QuestionSectionProps,
} from "./components/QuestionSection";

type ToMoInputData = {
  [key in MotivationKey]: number;
};

const App = () => {
  const [inputValue, setInputValue] = useState<ToMoInputData>({
    play: NaN,
    purpose: NaN,
    potential: NaN,
    emotionalPressure: NaN,
    economicPressure: NaN,
    inertia: NaN,
  });

  // useCallback しないと、ラジオを選択する度に再描画されて設問順序が入れ替わる
  const getHandleRadioChangeFn: QuestionSectionProps["curriedOnChangeFn"] =
    useCallback(
      (motivationKey: MotivationKey) =>
        (_: ChangeEvent<HTMLInputElement>, value: string) => {
          setInputValue((previous) => ({
            ...previous,
            [motivationKey]: Number(value),
          }));
        },
      []
    );

  /**
   * 診断するボタンを押した際の処理
   * ToMoを計算して表示する
   */
  const handleSubmit: FormEventHandler = useCallback(
    (event) => {
      event.preventDefault();
      console.log(inputValue);
    },
    [inputValue]
  );

  return (
    <Box>
      <form onSubmit={handleSubmit}>
        <QuestionSection
          tomoQuestions={tomoQuestions}
          curriedOnChangeFn={getHandleRadioChangeFn}
        />

        <Box sx={{ textAlign: "center" }}>
          <Button variant="contained" size="large" type="submit">
            診断する
          </Button>
        </Box>
      </form>
    </Box>
  );
};

export default App;
