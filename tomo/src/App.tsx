import Box from "@mui/material/Box";
import { tomoQuestions } from "./questions";
import { ToMoSurveyAnswer } from "./domain/tomo.type";
import RowRadioButtonsGroup from "./components/RadioGroup";

function App() {
  return (
    <Box sx={{ width: "100%", maxWidth: 1200, marginX: "200px" }}>
      {Object.keys(tomoQuestions)
        .toSorted(() => Math.random() - 0.5) // 考案者が本に書いた通り、質問文をランダムな順序で表示する
        .map((q) => {
          const question = q as keyof ToMoSurveyAnswer;
          return (
            <Box sx={{ marginBottom: "48px" }}>
              <RowRadioButtonsGroup formLabel={tomoQuestions[question].text} />
            </Box>
          );
        })}
    </Box>
  );
}

export default App;
