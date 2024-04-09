\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\ MY MY MY  /// error it stil sees values when it's gone over max///



WITH aa(x0,y0) as (select * from (SELECT ("values"->>'721f3b8c-1297-4c70-8b5c-202a0d0d7843')::numeric as high, ("values"->>'6877f6c9-96ac-450b-a9b4-f96bf94fe38c')::numeric as mass

                FROM public.directory_items

                where directoryid='615a677b-b5c3-451c-9a07-b3b3f04f52f5' order by high) t

                where high=trunc(1.7,1)),

                bb(x1,y1) as (select * from (SELECT ("values"->>'721f3b8c-1297-4c70-8b5c-202a0d0d7843')::numeric as high, ("values"->>'6877f6c9-96ac-450b-a9b4-f96bf94fe38c')::numeric as mass

                FROM public.directory_items

                where directoryid='615a677b-b5c3-451c-9a07-b3b3f04f52f5' order by high) t

                where   high=round(1.7,1))



select

                case

                               when 1.7=round(1.7,1) then aa.y0

                               when (bb.x1-aa.x0) = 0 then aa.y0

                else

                               (aa.y0 + (1.7-aa.x0)*((bb.y1-aa.y0)/(bb.x1-aa.x0)))

                end

                as field3

from aa, bb;





-- FUNCTION: public.alx_tst_intp(text, text, text, numeric)



-- DROP FUNCTION IF EXISTS public.alx_tst_intp(text, text, text, numeric);



CREATE OR REPLACE FUNCTION public.alx_tst_intp(

                id_high text,

                id_val text,

                id_dic text,

                val numeric)

    RETURNS numeric

    LANGUAGE 'plpgsql'

    COST 100

    VOLATILE PARALLEL UNSAFE

AS $BODY$

BEGIN

RETURN(

WITH aa(x0,y0) as (select * from (SELECT ("values"->>id_high)::numeric as high, ("values"->>id_val)::numeric as mass

                FROM public.directory_items

                where directoryid::text=id_dic order by high) t

                where high=trunc(val)),

                bb(x1,y1) as (select * from (SELECT ("values"->>id_high)::numeric as high, ("values"->>id_val)::numeric as mass

                FROM public.directory_items

                where directoryid::text=id_dic order by high) t

                where   high=round(val))



select

                case

                               when val=round(val) then aa.y0

                               when (bb.x1-aa.x0) = 0 then aa.y0

                else

                               (aa.y0 + (val-aa.x0)*((bb.y1-aa.y0)/(bb.x1-aa.x0)))

                end

                as field3

from aa, bb



);

END;

$BODY$;



ALTER FUNCTION public.alx_tst_intp(text, text, text, numeric)

    OWNER TO admin;







//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////NOT MY JUST MODIFIED FOR WORKING









-- FUNCTION: public.linear_interpolation(text, text, text, double precision)



-- DROP FUNCTION IF EXISTS public.linear_interpolation(text, text, text, double precision);



CREATE OR REPLACE FUNCTION public.linear_interpolation(

                id_high text,

                id_val text,

                id_dic text,

                val double precision)

    RETURNS double precision

    LANGUAGE 'plpgsql'

    COST 100

    VOLATILE PARALLEL UNSAFE

AS $BODY$

BEGIN

RETURN(



with

xy_table as (

                SELECT ("values" ->> id_high)::numeric as x,  ("values" ->> id_val)::numeric as y

                FROM public.directory_items

                where directoryid::text = id_dic order by x asc

),

pairs as (

select x, y from xy_table where x = (select min(x) as min from xy_table where x >= val)

                or x = (select max(x) as max from xy_table where x <= val)

),

x1 as (VALUES ((select x from pairs order by x asc limit 1))),

y1 as (VALUES ((select y from pairs order by x asc limit 1))),

x2 as (VALUES ((select x from pairs order by x desc limit 1))),

y2 as (VALUES ((select y from pairs order by x desc limit 1))),

k as (VALUES ((((table y2) - (table y1))/((table x2) -(table x1))))),

b as (VALUES (((table y1)-(table x1)*(table k))))

select

case

                when (select  count(distinct x) from pairs) > 1 then val*(table K)+(table b)

                else y

end  as res from pairs limit 1



)::double precision;

END;

$BODY$;



ALTER FUNCTION public.linear_interpolation(text, text, text, double precision)

    OWNER TO admin;







-- FUNCTION: public.linear_interpolation(text, text, text, double precision)



-- DROP FUNCTION IF EXISTS public.linear_interpolation(text, text, text, double precision);



CREATE OR REPLACE FUNCTION public.linear_interpolation(

                val double precision,

                id_high text,

                id_val text,

                id_dic text

                )

    RETURNS double precision

    LANGUAGE 'plpgsql'

    COST 100

    VOLATILE PARALLEL UNSAFE

AS $BODY$

BEGIN

RETURN(



with

xy_table as (

                SELECT ("values" ->> id_high)::numeric as x,  ("values" ->> id_val)::numeric as y

                FROM public.directory_items

                where directoryid::text = id_dic order by x asc

),

pairs as (

select x, y from xy_table where x = (select min(x) as min from xy_table where x >= val)

                or x = (select max(x) as max from xy_table where x <= val)

),

x1 as (VALUES ((select x from pairs order by x asc limit 1))),

y1 as (VALUES ((select y from pairs order by x asc limit 1))),

x2 as (VALUES ((select x from pairs order by x desc limit 1))),

y2 as (VALUES ((select y from pairs order by x desc limit 1))),

k as (VALUES ((((table y2) - (table y1))/((table x2) -(table x1))))),

b as (VALUES (((table y1)-(table x1)*(table k))))

select

case

                when (select  count(distinct x) from pairs) > 1 then val*(table K)+(table b)

                else y

end  as res from pairs limit 1



)::double precision;

END;

$BODY$;



ALTER FUNCTION public.linear_interpolation(text, text, text, double precision)

    OWNER TO admin;