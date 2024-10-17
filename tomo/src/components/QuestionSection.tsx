import Box from "@mui/material/Box";
import { memo } from "react";
import { type MotivationKey } from "../domain/tomo.type";
import { type ToMoQuestion } from "../questions";
import { type RadioGroupOnChangeFn, RowRadioButtonsGroup } from "./RadioGroup";

export type QuestionSectionProps = {
  tomoQuestions: ToMoQuestion;
  curriedOnChangeFn: (motivationKey: MotivationKey) => RadioGroupOnChangeFn;
};

/**
 * ToMoを計測する設問を表示するコンポーネント
 */
export const QuestionSection = memo(
  ({ tomoQuestions, curriedOnChangeFn }: QuestionSectionProps) => {
    return (
      <Box>
        {Object.keys(tomoQuestions)
          .toSorted(() => Math.random() - 0.5) // 考案者が本に書いた通り、質問文をランダムな順序で表示する
          .map((q, i) => {
            const motivationKey = q as MotivationKey;
            const handleRadioChange = curriedOnChangeFn(motivationKey);
            return (
              <Box key={i} sx={{ marginBottom: "48px" }}>
                <RowRadioButtonsGroup
                  formLabel={tomoQuestions[motivationKey].text}
                  onChangeFn={handleRadioChange}
                />
              </Box>
            );
          })}
      </Box>
    );
  }
);
