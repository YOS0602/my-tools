import { memo } from "react";
import Radio from "@mui/material/Radio";
import RadioGroup from "@mui/material/RadioGroup";
import FormControlLabel from "@mui/material/FormControlLabel";
import FormLabel from "@mui/material/FormLabel";
import FormControl from "@mui/material/FormControl";

type RowRadioButtonsGroupProps = {
  formLabel: string;
};

export const RowRadioButtonsGroup = memo(
  ({ formLabel }: RowRadioButtonsGroupProps) => {
    return (
      <FormControl>
        {/* FIXME id や name を正しい値にする */}
        <FormLabel id="demo-row-radio-buttons-group-label">
          {formLabel}
        </FormLabel>
        <RadioGroup
          row
          aria-labelledby="demo-row-radio-buttons-group-label"
          name="row-radio-buttons-group"
        >
          <FormControlLabel
            value={1}
            control={<Radio />}
            label="まったく違う"
            labelPlacement="bottom"
            // TODO スマホサイズの時は24か指定なしにしたい
            sx={{ "& .MuiSvgIcon-root": { fontSize: 64 } }}
          />
          <FormControlLabel
            value={2}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 64 } }}
          />
          <FormControlLabel
            value={3}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 64 } }}
          />
          <FormControlLabel
            value={4}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 64 } }}
          />
          <FormControlLabel
            value={5}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 64 } }}
          />
          <FormControlLabel
            value={6}
            control={<Radio />}
            label=""
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 64 } }}
          />
          <FormControlLabel
            value={7}
            control={<Radio />}
            label="まったくその通りである"
            labelPlacement="bottom"
            sx={{ "& .MuiSvgIcon-root": { fontSize: 64 } }}
          />
        </RadioGroup>
      </FormControl>
    );
  }
);
