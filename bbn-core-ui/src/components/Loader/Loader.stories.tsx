import type { Meta, StoryObj } from "@storybook/react";

import { Loader } from "./Loader";

const meta: Meta<typeof Loader> = {
  component: Loader,
  tags: ["autodocs"],
};

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    className: "text-accent-secondary",
  },
};
