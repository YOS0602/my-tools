import Radio from "@mui/material/Radio";
import RadioGroup from "@mui/material/RadioGroup";
import FormControlLabel from "@mui/material/FormControlLabel";
import FormControl from "@mui/material/FormControl";
import FormLabel from "@mui/material/FormLabel";

type RowRadioButtonsGroupProps = {
  formLabel: string;
};

export default function RowRadioButtonsGroup({
  formLabel,
}: RowRadioButtonsGroupProps) {
  return (
    <FormControl>
      <FormLabel id="demo-row-radio-buttons-group-label">{formLabel}</FormLabel>
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
        />
        <FormControlLabel
          value={2}
          control={<Radio />}
          label=""
          labelPlacement="bottom"
        />
        <FormControlLabel
          value={3}
          control={<Radio />}
          label=""
          labelPlacement="bottom"
        />
        <FormControlLabel
          value={4}
          control={<Radio />}
          label=""
          labelPlacement="bottom"
        />
        <FormControlLabel
          value={5}
          control={<Radio />}
          label=""
          labelPlacement="bottom"
        />
        <FormControlLabel
          value={6}
          control={<Radio />}
          label=""
          labelPlacement="bottom"
        />
        <FormControlLabel
          value={7}
          control={<Radio />}
          label="まったくその通りである"
          labelPlacement="bottom"
        />
      </RadioGroup>
    </FormControl>
  );
}
