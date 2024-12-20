import {
  type ChangeEvent,
  type FormEventHandler,
  useCallback,
  useState,
} from "react";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { tomoQuestions } from "./libs/tomoQuestions";
import { ToMoSurveyAnswer, type MotivationKey } from "./libs/tomo.type";
import { calculateToMo } from "./libs/tomoCalculator";
import {
  QuestionSection,
  type QuestionSectionProps,
} from "./components/QuestionSection";

const App = () => {
  const [inputValue, setInputValue] = useState<ToMoSurveyAnswer>({
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

      // TODO バリデーションして NaN が一つでもあったら入力を促すメッセージを表示する
      // それか初期値として「4」を選択させてstateにもsetしておくなど検討
      const tomoScore = calculateToMo(inputValue);
      // TODO alertではなくModalなどを使って表示する
      alert(
        isNaN(tomoScore)
          ? "ToMo指数の計算に失敗しました！"
          : `あなたのToMo指数は...... ${tomoScore}点です！`
      );
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
