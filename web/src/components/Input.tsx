interface InputProps {
  label: string;
  id: string;
  placeholder: string;
  type: string;
}

const Input = ({ label, id, placeholder, type }: InputProps) => {
  return (
    <div className="relative p-3 bg-[#253029] rounded-xl border border-[#3c5045]/50">
      <label htmlFor={id} className="text-[#a0b9a6] text-xs font-medium">
        {label}
      </label>
      <input
        id={id}
        type={type}
        className="w-full bg-transparent focus:outline-none text-[#d3e4cd] placeholder:text-[#a0b9a6]/50"
        placeholder={placeholder}
      />
    </div>
  );
};

export default Input;
