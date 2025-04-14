interface InputProps {
  label: string;
  id: string;
  placeholder: string;
  type: string;
  onChange: (char: string) => void;
  value: string;
}

const Input = ({
  label,
  id,
  placeholder,
  type,
  onChange,
  value,
}: InputProps) => {
  return (
    <div className="relative p-2 sm:p-3 bg-[#253029] rounded-xl border border-[#3c5045]/50">
      <label
        htmlFor={id}
        className="text-[#a0b9a6] text-[10px] sm:text-xs font-medium"
      >
        {label}
      </label>
      <input
        value={value}
        onChange={(e) => onChange(e.target.value)}
        id={id}
        type={type}
        className="w-full bg-transparent focus:outline-none text-[#d3e4cd] placeholder:text-[#a0b9a6]/50 text-sm sm:text-base"
        placeholder={placeholder}
      />
    </div>
  );
};

export default Input;
