import { memo } from "react";
import Radio from "@mui/material/Radio";
import RadioGroup from "@mui/material/RadioGroup";
import FormControlLabel from "@mui/material/FormControlLabel";
import FormLabel from "@mui/material/FormLabel";
import FormControl from "@mui/material/FormControl";

type RowRadioButtonsGroupProps = {
  formLabel: string;
  onChangeFn: RadioGroupOnChangeFn;
};

export type RadioGroupOnChangeFn = (
  event: React.ChangeEvent<HTMLInputElement>,
  value: string
) => void;

export const RowRadioButtonsGroup = memo(
  ({ formLabel, onChangeFn }: RowRadioButtonsGroupProps) => {
    return (
      <FormControl>
        <FormLabel>{formLabel}</FormLabel>
        <RadioGroup row onChange={onChangeFn}>
          <FormControlLabel
            value={"1"}
            control={<Radio />}
            label="まったく違う"
            labelPlacement="bottom"
            // TODO スマホサイズの時は24か指定なしにしたい
            sx={{ "& .MuiSvgIcon-root": { fontSize: 48 } }}
          />
          <FormControlLabel
            value={"2"}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 48 } }}
          />
          <FormControlLabel
            value={"3"}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 48 } }}
          />
          <FormControlLabel
            value={"4"}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 48 } }}
          />
          <FormControlLabel
            value={"5"}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 48 } }}
          />
          <FormControlLabel
            value={"6"}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 48 } }}
          />
          <FormControlLabel
            value={"7"}
            control={<Radio />}
            label="まったくその通りである"
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 48 } }}
          />
        </RadioGroup>
      </FormControl>
    );
  }
);
