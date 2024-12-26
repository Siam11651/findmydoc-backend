create or replace function get_distance(param_point0 point, param_point1 point)
returns float8
language plpgsql
as $$
	declare
		distance float8;
	begin
		select ST_DistanceSphere(
		    public.ST_MakePoint(param_point0[0], param_point0[1]),
		    public.ST_MakePoint(param_point1[0], param_point1[1])
		) / 1000 into distance;

		return distance;
	end;
$$;