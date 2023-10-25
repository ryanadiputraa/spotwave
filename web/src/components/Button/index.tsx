import { useEffect, useState } from 'react';

type Variants = 'primary' | 'secondary';

interface Props {
	children: React.ReactNode;
	variant?: Variants;
	classNames?: string;
	onClick?: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => any;
}

export const Button = ({ children, variant = 'primary', classNames, onClick = () => {} }: Props) => {
	const [classVariants, setClassVariants] = useState('');

	useEffect(() => {
		switch (variant) {
			case 'primary':
				setClassVariants('bg-black text-white');
				break;

			case 'secondary':
				setClassVariants('bg-white text-black');
				break;

			default:
				break;
		}
	}, []);

	return (
		<button
			className={`px-4 py-2 rounded-md active:scale-95 ${classVariants} ${classNames}`}
			onClick={(e) => onClick(e)}
		>
			{children}
		</button>
	);
};
