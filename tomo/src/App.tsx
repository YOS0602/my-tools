import { FormEventHandler } from "react";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { tomoQuestions } from "./questions";
import { ToMoSurveyAnswer } from "./domain/tomo.type";
import { RowRadioButtonsGroup } from "./components/RadioGroup";

const App = () => {
  const onSubmit: FormEventHandler = (event) => {
    event.preventDefault();
    console.log(event);
  };

  return (
    <Box>
      <form onSubmit={onSubmit}>
        {Object.keys(tomoQuestions)
          .toSorted(() => Math.random() - 0.5) // 考案者が本に書いた通り、質問文をランダムな順序で表示する
          .map((q, i) => {
            const question = q as keyof ToMoSurveyAnswer;
            return (
              <Box key={i} sx={{ marginBottom: "48px" }}>
                <RowRadioButtonsGroup
                  formLabel={tomoQuestions[question].text}
                />
              </Box>
            );
          })}

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
