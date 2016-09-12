/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   problem5.c                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: ale-batt <ale-batt@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/07 16:20:40 by ale-batt          #+#    #+#             */
/*   Updated: 2016/09/07 16:25:32 by ale-batt         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <stdio.h>

int is_smallest_multiple(int nb)
{
	int max = 20;
	int n = 2;

	while (n < 20)
	{
		if ((nb % n) != 0)
			return 0;
		n++;
	}
	return 1;
}

int main(void)
{
	int i = 10;

	while (42)
	{
		if (is_smallest_multiple(i))
			break ;
		i++;
	}
	printf("Result = %d \n", i);
	return (0);
}
